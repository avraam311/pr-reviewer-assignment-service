package statistics

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
)

type Repository interface {
	GetStatistics(context.Context, string) (*db.Statistics, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
