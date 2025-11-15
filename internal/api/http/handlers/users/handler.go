package users

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

type Service interface {
	SetIsActive(context.Context, *dto.UserWithIsActive) (*db.User, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
