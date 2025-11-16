package teams

import (
	"context"
	"fmt"
)

func (s *Service) DeactivateTeamUsers(ctx context.Context, teamName string) error {
	err := s.repo.DeactivateTeamUsers(ctx, teamName)
	if err != nil {
		return fmt.Errorf("service/deactivate_team_users.go - %w", err)
	}

	return nil
}
