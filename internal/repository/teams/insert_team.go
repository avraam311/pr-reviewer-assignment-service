package teams

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) InsertTeam(ctx context.Context, team *dto.Team) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("repository/insert_team.go - failed to begin transaction - %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	queryTeam := `
        INSERT INTO team (team_name)
        VALUES ($1)
    `
	_, err = tx.Exec(ctx, queryTeam, team.TeamName)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrTeamAlreadyExists
		}

		return fmt.Errorf("repository/insert_team.go - failed to insert team - %w", err)
	}

	queryUser := `
    INSERT INTO "user" (user_id, username, is_active, team_name)
    VALUES ($1, $2, $3, $4)
    ON CONFLICT (user_id) DO UPDATE 
    SET username = EXCLUDED.username,
        is_active = EXCLUDED.is_active,
        team_name = EXCLUDED.team_name
`
	for _, m := range team.Members {
		_, err = tx.Exec(ctx, queryUser, m.UserID, m.UserName, m.IsActive, team.TeamName)
		if err != nil {
			return fmt.Errorf("repository/insert_team.go - failed to insert or update user - %w", err)
		}
	}

	return nil
}
