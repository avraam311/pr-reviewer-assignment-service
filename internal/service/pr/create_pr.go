package pr

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (s *Service) CreatePR(ctx context.Context, pr *dto.PR) (*db.PR, error) {
	prDB, err := s.repo.CreatePR(ctx, pr)
	if err != nil {
		return nil, fmt.Errorf("service/create_pr.go - %w", err)
	}

	return prDB, nil
}
