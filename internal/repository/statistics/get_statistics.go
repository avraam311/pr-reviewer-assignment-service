package statistics

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
)

func (r *Repository) GetStatistics(ctx context.Context, userID string) (*db.Statistics, error) {
	queryExists := `
        SELECT EXISTS(SELECT 1 FROM "user" WHERE user_id = $1)
    `
	var exists bool
	err := r.db.QueryRow(ctx, queryExists, userID).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("repository/get_statistics.go - failed to check if user exists - %w", err)
	}
	if !exists {
		return nil, ErrUserNotFound
	}
	query := `
        SELECT COUNT(*)
        FROM pull_request
        WHERE $1 = ANY(assigned_reviewers)
    `
	var assignedCount int
	err = r.db.QueryRow(ctx, query, userID).Scan(&assignedCount)
	if err != nil {
		return nil, fmt.Errorf("repository/get_statistics.go - failed to scan assigned count - %w", err)
	}

	stats := &db.Statistics{
		UserID:          userID,
		AssignmentCount: assignedCount,
	}
	return stats, nil
}
