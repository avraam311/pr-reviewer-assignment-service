package teams

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

type Repository interface {
	InsertTeam(context.Context, *dto.Team) error
	SelectTeam(context.Context, string) (*dto.Team, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
