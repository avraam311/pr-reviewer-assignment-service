package teams

import (
	"context"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
)

func (r *Repository) SelectTeam(ctx context.Context, teamName string) (*dto.Team, error) {
	query := `
        SELECT t.team_name, u.user_id, u.username, u.is_active
        FROM team t
        LEFT JOIN user u ON t.team_name = u.team_name
        WHERE t.team_name = $1
    `
	rows, err := r.db.Query(ctx, query, teamName)
	if err != nil {
		return nil, fmt.Errorf("repository/select_team.go - failed to query team - %w", err)
	}
	defer rows.Close()

	var team dto.Team
	team.TeamName = teamName
	team.Members = make([]*dto.User, 0)

	for rows.Next() {
		var user dto.User
		var teamName string

		err := rows.Scan(&teamName, &user.UserID, &user.UserName, &user.IsActive)
		if err != nil {
			return nil, fmt.Errorf("repository/select_team - failed to scan row - %w", err)
		}

		if user.UserID != "" {
			team.Members = append(team.Members, &user)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("repository/select_team - rows iteration error - %w", err)
	}

	return &team, nil
}
