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
		return nil, fmt.Errorf("%w: failed to get status of PR: %v", ErrPROrOldUserNotFound, err)
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
		return nil, fmt.Errorf("%w: failed to get team of old reviewer: %v", ErrPROrOldUserNotFound, err)
	}

	queryGetRandomActiveReviewer := `
        SELECT user_id FROM "user"
        WHERE team_name = $1 AND user_id <> $2 AND is_active = true
        ORDER BY random()
        LIMIT 1
    `
	var newReviewer string
	err = r.db.QueryRow(ctx, queryGetRandomActiveReviewer, teamName, pr.OldUserID).Scan(&newReviewer)
	if err == pgx.ErrNoRows {
		return nil, fmt.Errorf("no replacement reviewer found in team")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to select replacement reviewer: %w", err)
	}

	queryGetAssignedReviewers := `
        SELECT assigned_reviewers 
		FROM pull_request 
		WHERE pull_request_id = $1
    `
	var assignedReviewers []string
	err = r.db.QueryRow(ctx, queryGetAssignedReviewers, pr.PRID).Scan(&assignedReviewers)
	if err != nil {
		return nil, fmt.Errorf("failed to get current assigned reviewers: %w", err)
	}

	replaced := false
	for i, id := range assignedReviewers {
		if id == pr.OldUserID {
			assignedReviewers[i] = newReviewer
			replaced = true
			break
		}
	}
	if !replaced {
		return nil, ErrPROrOldUserNotFound
	}

	queryUpdateAssignedReviewers := `
        UPDATE pull_request
        SET assigned_reviewers = $1
        WHERE pull_request_id = $2
        RETURNING pull_request_id, pull_request_name, author_id, status, assigned_reviewers
    `
	var updatedPR db.PRWithReplacedBy
	var prData db.PR
	err = r.db.QueryRow(ctx, queryUpdateAssignedReviewers, assignedReviewers, pr.PRID).Scan(
		&prData.PRID,
		&prData.PRName,
		&prData.AuthorID,
		&prData.Status,
		&prData.AssignedReviewers,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to update assigned reviewers: %w", err)
	}
	updatedPR.PR = &prData
	updatedPR.ReplacedBy = newReviewer

	return &updatedPR, nil
}
