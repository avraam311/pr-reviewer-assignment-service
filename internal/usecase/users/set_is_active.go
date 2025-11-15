package users

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (s *Service) SetIsActive(ctx context.Context, usr *dto.UserWithIsActive) (*db.User, error) {
	user, err := s.repo.UpdateUser(ctx, usr)
	if err != nil {
		return nil, fmt.Errorf("service/set_is_active.go - %w", err)
	}

	return user, nil
}
