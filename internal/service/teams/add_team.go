package teams

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (s *Service) AddTeam(ctx context.Context, team *dto.Team) error {
	err := s.repo.CreateTeam(ctx, team)
	if err != nil {
		return fmt.Errorf("service/add_team.go - %w", err)
	}

	return nil
}
