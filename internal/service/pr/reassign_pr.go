package pr

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (s *Service) ReassignPR(ctx context.Context, pr *dto.PRWithOldUserID) (*db.PRWithReplacedBy, error) {
	prDB, err := s.repo.ReassignPRReviewer(ctx, pr)
	if err != nil {
		return nil, fmt.Errorf("service/reassign_pr.go - %w", err)
	}

	return prDB, nil
}
