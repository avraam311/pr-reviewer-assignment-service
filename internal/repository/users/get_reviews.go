package users

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
)

func (r *Repository) GetReviews(ctx context.Context, userID string) ([]*db.PRShort, error) {
	query := `
        SELECT pull_request_id, pull_request_name, author_id, status
        FROM pull_request
        WHERE $1 = ANY(assigned_reviewers)
    `
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("repository/get_reviews.go - failed to query reviews - %w", err)
	}
	defer rows.Close()

	var prs []*db.PRShort
	for rows.Next() {
		var pr db.PRShort
		if err := rows.Scan(&pr.PRID, &pr.PRName, &pr.AuthorID, &pr.Status); err != nil {
			return nil, fmt.Errorf("repository/get_reviews.go - failed to scan pr - %w", err)
		}
		prs = append(prs, &pr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/get_reviews.go - rows iteration error - %w", err)
	}

	return prs, nil
}
