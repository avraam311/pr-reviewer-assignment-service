package teams

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

type Repository interface {
	CreateTeam(context.Context, *dto.Team) error
	GetTeam(context.Context, string) (*dto.Team, error)
	DeactivateTeamUsers(context.Context, string) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
