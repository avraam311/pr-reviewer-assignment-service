package db

type User struct {
	UserID   string `json:"user_id" db:"user_id" binding:"required"`
	UserName string `json:"username" db:"username" binding:"required"`
	TeamName string `json:"team_name" db:"team_name" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active" binding:"required"`
}
