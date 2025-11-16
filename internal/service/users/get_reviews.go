package users

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
)

func (s *Service) GetReviews(ctx context.Context, userID string) ([]*db.PRShort, error) {
	reviews, err := s.repo.GetReviews(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("service/get_reviews.go - %w", err)
	}

	return reviews, nil
}
