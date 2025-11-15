package teams

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (s *Service) GetTeam(ctx context.Context, teamName string) (*dto.Team, error) {
	team, err := s.repo.SelectTeam(ctx, teamName)
	if err != nil {
		return nil, fmt.Errorf("service/get_team.go - %w", err)
	}

	return team, nil
}
