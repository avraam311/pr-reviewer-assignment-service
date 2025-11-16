package statistics

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
)

func (s *Service) GetStatistics(ctx context.Context, userID string) (*db.Statistics, error) {
	stats, err := s.repo.GetStatistics(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("service/get_statistics.go - %w", err)
	}

	return stats, nil
}
