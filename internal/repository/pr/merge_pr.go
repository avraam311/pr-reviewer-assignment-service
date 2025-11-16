package pr

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) MergePR(ctx context.Context, pr *dto.PRWithPRID) (*db.PRWithMergedAt, error) {
	query := `
        UPDATE pull_request
        SET status = $1, merged_at = COALESCE(merged_at, NOW())
        WHERE pull_request_id = $2 AND status != $1
        RETURNING pull_request_id, pull_request_name, author_id, status, assigned_reviewers, merged_at
    `
	var updatedPR db.PRWithMergedAt
	err := r.db.QueryRow(ctx, query, prStatusMerged, pr.PRID).Scan(
		&updatedPR.PRID,
		&updatedPR.PRName,
		&updatedPR.AuthorID,
		&updatedPR.Status,
		&updatedPR.AssignedReviewers,
		&updatedPR.MergedAt,
	)
	if err == pgx.ErrNoRows {
		selectQuery := `
            SELECT pull_request_id, pull_request_name, author_id, status, assigned_reviewers, merged_at
            FROM pull_request
            WHERE pull_request_id = $1
        `
		err = r.db.QueryRow(ctx, selectQuery, pr.PRID).Scan(
			&updatedPR.PRID,
			&updatedPR.PRName,
			&updatedPR.AuthorID,
			&updatedPR.Status,
			&updatedPR.AssignedReviewers,
			&updatedPR.MergedAt,
		)
		if err == pgx.ErrNoRows {
			return nil, ErrPRNotFound
		}
		if err != nil {
			return nil, fmt.Errorf("repository/merge_pr.go - failed to fetch pull request after merge conflict: %w", err)
		}
		return &updatedPR, nil
	}
	if err != nil {
		return nil, fmt.Errorf("repository/merge_pr.go - failed to merge pull request: %w", err)
	}
	return &updatedPR, nil
}
