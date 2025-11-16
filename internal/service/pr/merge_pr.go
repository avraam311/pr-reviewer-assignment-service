package pr

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (s *Service) MergePR(ctx context.Context, pr *dto.PRWithPRID) (*db.PRWithMergedAt, error) {
	prDB, err := s.repo.MergePR(ctx, pr)
	if err != nil {
		return nil, fmt.Errorf("service/merge_pr.go - %w", err)
	}

	return prDB, nil
}
