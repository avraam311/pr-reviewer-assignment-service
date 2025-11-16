package users

import (
	"context"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/db"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) UpdateUserIsActive(ctx context.Context, usr *dto.UserWithIsActive) (*db.User, error) {
	query := `
        UPDATE "user"
        SET is_active = $1
        WHERE user_id = $2
        RETURNING user_id, username, team_name, is_active
    `

	var user db.User
	err := r.db.QueryRow(ctx, query, usr.IsActive, usr.UserID).Scan(
		&user.UserID, &user.UserName, &user.TeamName, &user.IsActive,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}
