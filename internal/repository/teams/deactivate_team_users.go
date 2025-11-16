package teams

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r *Repository) DeactivateTeamUsers(ctx context.Context, teamName string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("repository/deactivate_team_users.go - failed to begin transaction - %w", err)
	}
	defer tx.Rollback(ctx)

	queryTeamUpdate := `
		UPDATE "user" SET is_active = false
        WHERE team_name = $1 AND is_active = True
	`
	_, err = tx.Exec(ctx, queryTeamUpdate, teamName)
	if err != nil {
		return fmt.Errorf("repository/deactivate_team_users.go - failed to deactivate users - %w", err)
	}

	querySelectPRAndUsers := `
		SELECT pr.pull_request_id, unr.user_id
		FROM pull_request pr
		JOIN LATERAL unnest(pr.assigned_reviewers) AS unr(user_id) ON TRUE
		JOIN "user" u ON u.user_id = unr.user_id
		WHERE pr.status = 'OPEN' AND u.is_active = FALSE AND u.team_name = $1
	`
	rows, err := tx.Query(ctx, querySelectPRAndUsers, teamName)
	if err != nil {
		return fmt.Errorf("repository/deactivate_team_users.go - failed to get prs and users - %w", err)
	}
	defer rows.Close()

	type prReviewer struct {
		prID   string
		userID string
	}

	var toReplace []prReviewer
	for rows.Next() {
		var prID, userID string
		if err := rows.Scan(&prID, &userID); err != nil {
			return fmt.Errorf("repository/deactivate_team_users.go - failed to scan pr or user - %w", err)
		}
		toReplace = append(toReplace, prReviewer{prID, userID})
	}

	querySelectID := `
		SELECT user_id 
		FROM "user"
		WHERE team_name = $1
		AND is_active = TRUE
		AND user_id <> $3
		AND user_id <> (SELECT author_id FROM pull_request WHERE pull_request_id = $2)
		AND user_id NOT IN (
			SELECT unnest(assigned_reviewers) FROM pull_request WHERE pull_request_id = $2
		)
		ORDER BY random() LIMIT 1
	`
	for _, prRev := range toReplace {
		var newReviewerID string
		err = tx.QueryRow(ctx, querySelectID, teamName, prRev.prID, prRev.userID).Scan(&newReviewerID)

		if err == pgx.ErrNoRows {
			queryDeleteReviewer := `
				UPDATE pull_request
				SET assigned_reviewers = array_remove(assigned_reviewers, $2)
				WHERE pull_request_id = $1
        	`
			_, err = tx.Exec(ctx, queryDeleteReviewer, prRev.prID, prRev.userID)
			if err != nil {
				return fmt.Errorf("repository/deactivate_team_users.go - failed to remove reviewer - %w", err)
			}
		} else if err != nil {
			return fmt.Errorf("repository/deactivate_team_users.go - failed to select new reviewer - %w", err)
		} else {
			queryUpdateReviewer := `
				UPDATE pull_request
				SET assigned_reviewers = array_replace(assigned_reviewers, $2, $3)
				WHERE pull_request_id = $1
        	`
			_, err = tx.Exec(ctx, queryUpdateReviewer, prRev.prID, prRev.userID, newReviewerID)
			if err != nil {
				return fmt.Errorf("repository/deactivate_team_users.go - failed to replace reviewer - %w", err)
			}
		}
	}

	return tx.Commit(ctx)
}
