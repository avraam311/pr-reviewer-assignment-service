package teams

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (r *Repository) GetTeam(ctx context.Context, teamName string) (*dto.Team, error) {
	query := `
        SELECT t.team_name, u.user_id, u.username, u.is_active
        FROM team t
        LEFT JOIN "user" u ON t.team_name = u.team_name
        WHERE t.team_name = $1
    `
	rows, err := r.db.Query(ctx, query, teamName)
	if err != nil {
		return nil, fmt.Errorf("repository/select_team.go - failed to query team - %w", err)
	}
	defer rows.Close()

	var team dto.Team
	team.Members = make([]*dto.User, 0)

	found := false
	for rows.Next() {
		found = true
		var user dto.User
		var tName string

		err := rows.Scan(&tName, &user.UserID, &user.UserName, &user.IsActive)
		if err != nil {
			return nil, fmt.Errorf("repository/select_team - failed to scan row - %w", err)
		}

		team.TeamName = tName

		if user.UserID != "" {
			team.Members = append(team.Members, &user)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/select_team - rows iteration error - %w", err)
	}

	if !found {
		return nil, ErrTeamNotFound
	}

	return &team, nil
}
