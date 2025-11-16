package statistics

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
)

type Service interface {
	GetStatistics(context.Context, string) (*db.Statistics, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
