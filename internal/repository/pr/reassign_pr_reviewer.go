package pr

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) ReassignPRReviewer(ctx context.Context, pr *dto.PRWithOldUserID) (*db.PRWithReplacedBy, error) {
	queryGetPRStatus := `
        SELECT status 
        FROM pull_request 
        WHERE pull_request_id = $1
    `
	var status string
	err := r.db.QueryRow(ctx, queryGetPRStatus, pr.PRID).Scan(&status)
	if err != nil {
		return nil, ErrPROrOldUserNotFound
	}
	if status == prStatusMerged {
		return nil, ErrReassignAfterMerge
	}

	queryGetTeamName := `
        SELECT team_name 
        FROM "user" 
        WHERE user_id = $1 AND is_active = true
    `
	var teamName string
	err = r.db.QueryRow(ctx, queryGetTeamName, pr.OldUserID).Scan(&teamName)
	if err != nil {
		return nil, ErrPROrOldUserNotFound
	}

	queryCheckReviewerAssigned := `
        SELECT 1 FROM pull_request 
        WHERE pull_request_id = $1 AND $2 = ANY(assigned_reviewers)
    `
	var dummy int
	err = r.db.QueryRow(ctx, queryCheckReviewerAssigned, pr.PRID, pr.OldUserID).Scan(&dummy)
	if err == pgx.ErrNoRows {
		return nil, ErrReviewerNotAssigned
	}
	if err != nil {
		return nil, fmt.Errorf("repository/reassign_pr_reviewer - failed to check if reviewer assigned - %w", err)
	}

	queryAuthorID := `
        SELECT author_id 
        FROM pull_request 
        WHERE pull_request_id = $1  
    `
	var authorID string
	err = r.db.QueryRow(ctx, queryAuthorID, pr.PRID).Scan(&authorID)
	if err != nil {
		return nil, fmt.Errorf("repository/reassign_pr_reviewer - failed to get author_id - %w", err)
	}

	queryGetRandomActiveReviewer := `
        SELECT user_id FROM "user"
        WHERE team_name = $1 AND user_id <> $2 AND user_id <> $3 AND is_active = true
          AND user_id != ALL (
            SELECT unnest(assigned_reviewers) FROM pull_request WHERE pull_request_id = $4
          )
        ORDER BY random()
        LIMIT 1
    `
	var newReviewer string
	err = r.db.QueryRow(ctx, queryGetRandomActiveReviewer, teamName, pr.OldUserID, authorID, pr.PRID).Scan(&newReviewer)
	if err == pgx.ErrNoRows {
		// Нет кандидата на замену, возвращаем ошибку и оставляем старого ревьюера
		var prData db.PR
		selectQuery := `
            SELECT pull_request_id, pull_request_name, author_id, status, assigned_reviewers
            FROM pull_request
            WHERE pull_request_id = $1
        `
		err = r.db.QueryRow(ctx, selectQuery, pr.PRID).Scan(
			&prData.PRID, &prData.PRName, &prData.AuthorID, &prData.Status, &prData.AssignedReviewers,
		)
		if err != nil {
			return nil, fmt.Errorf("repository/reassign_pr_reviewer - failed to fetch pull request when no candidate found - %w", err)
		}
		return &db.PRWithReplacedBy{
			PR:         &prData,
			ReplacedBy: pr.OldUserID,
		}, ErrNoCandidate
	}
	if err != nil {
		return nil, fmt.Errorf("repository/reassign_pr_reviewer - failed to select replacement reviewer - %w", err)
	}

	// Выполняем замену старого ревьюера на нового, уникальность уже гарантирована по запросу
	queryUpdateAssignedReviewers := `
        UPDATE pull_request
        SET assigned_reviewers = ARRAY_REPLACE(assigned_reviewers, $1, $2)
        WHERE pull_request_id = $3
        RETURNING pull_request_id, pull_request_name, author_id, status, assigned_reviewers
    `
	var updatedPR db.PRWithReplacedBy
	var prData db.PR
	err = r.db.QueryRow(ctx, queryUpdateAssignedReviewers, pr.OldUserID, newReviewer, pr.PRID).Scan(
		&prData.PRID,
		&prData.PRName,
		&prData.AuthorID,
		&prData.Status,
		&prData.AssignedReviewers,
	)
	if err != nil {
		return nil, fmt.Errorf("repository/reassign_pr_reviewer - failed to update assigned reviewers - %w", err)
	}
	updatedPR.PR = &prData
	updatedPR.ReplacedBy = newReviewer

	return &updatedPR, nil
}
