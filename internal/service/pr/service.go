package pr

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

type Repository interface {
	CreatePR(context.Context, *dto.PR) (*db.PR, error)
	MergePR(context.Context, *dto.PRWithPRID) (*db.PRWithMergedAt, error)
	ReassignPRReviewer(context.Context, *dto.PRWithOldUserID) (*db.PRWithReplacedBy, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
