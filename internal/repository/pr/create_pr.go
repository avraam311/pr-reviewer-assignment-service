package pr

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

const (
	prStatusOpen     = "OPEN"
	userIsActiveTrue = true
)

func (r *Repository) CreatePR(ctx context.Context, pr *dto.PR) (*db.PR, error) {
	queryTeamName := `
		SELECT team_name 
		FROM "user" 
		WHERE user_id = $1
	`
	var teamName string
	err := r.db.QueryRow(ctx, queryTeamName, pr.AuthorID).Scan(&teamName)
	if err == pgx.ErrNoRows {
		return nil, ErrAuthorNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("repository/create_pr.go - failed to get author team_name: %w", err)
	}

	queryCheckPR := `
		SELECT 1 
		FROM pull_request 
		WHERE pull_request_id = $1
	`
	var exists int
	err = r.db.QueryRow(ctx, queryCheckPR, pr.PRID).Scan(&exists)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("repository/create_pr.go - failed to check pr existence: %w", err)
	}
	if exists == 1 {
		return nil, ErrPRAlreadyExists
	}

	queryReviewers := `
        SELECT user_id
        FROM "user"
        WHERE team_name = $1 AND user_id <> $2 AND is_active = $3
        ORDER BY random()
        LIMIT 2
    `
	rows, err := r.db.Query(ctx, queryReviewers, teamName, pr.AuthorID, userIsActiveTrue)
	if err != nil {
		return nil, fmt.Errorf("repository/create_pr.go - failed to get reviewers: %w", err)
	}
	defer rows.Close()

	assigned := make([]string, 0, 2)
	for rows.Next() {
		var reviewerID string
		if err := rows.Scan(&reviewerID); err != nil {
			return nil, fmt.Errorf("repository/create_pr.go - failed to scan reviewer id: %w", err)
		}
		assigned = append(assigned, reviewerID)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/create_pr.go - error iterating reviewers: %w", err)
	}

	queryPR := `
        INSERT INTO pull_request (pull_request_id, pull_request_name, author_id, status, assigned_reviewers)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING pull_request_id, pull_request_name, author_id, status, assigned_reviewers
    `
	var createdPR db.PR
	err = r.db.QueryRow(ctx, queryPR, pr.PRID, pr.PRName, pr.AuthorID, prStatusOpen, assigned).Scan(
		&createdPR.PRID, &createdPR.PRName, &createdPR.AuthorID, &createdPR.Status, &createdPR.AssignedReviewers,
	)
	if err != nil {
		return nil, fmt.Errorf("repository/create_pr.go - failed to insert pull request: %w", err)
	}

	return &createdPR, nil
}
