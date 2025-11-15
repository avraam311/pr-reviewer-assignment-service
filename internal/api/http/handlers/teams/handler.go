package teams

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"

	"github.com/go-playground/validator/v10"
)

type Service interface {
	AddTeam(context.Context, *dto.Team) (error)
	GetTeam(context.Context, string) (*dto.Team, error)
}

type Handler struct {
	service   Service
	validator *validator.Validate
}

func NewHandler(service Service, validator *validator.Validate) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}
