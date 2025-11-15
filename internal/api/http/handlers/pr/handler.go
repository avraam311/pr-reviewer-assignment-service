package pr

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

type Service interface {
	CreatePR(context.Context, *dto.PR) (*db.PR, error)
	MergePR(context.Context, *dto.PRWithPRID) (*db.PRWithMergedAt, error)
	ReassignPR(context.Context, *dto.PRWithOldUserID) (*db.PRWithReplacedBy, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}
