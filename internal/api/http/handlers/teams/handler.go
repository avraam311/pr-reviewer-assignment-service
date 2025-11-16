package teams

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

type Service interface {
	AddTeam(context.Context, *dto.Team) error
	GetTeam(context.Context, string) (*dto.Team, error)
	DeactivateTeamUsers(context.Context, string) error
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
