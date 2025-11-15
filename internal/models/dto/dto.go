package dto

type Team struct {
	TeamName string    `json:"team_name" db:"team_name" binding:"required"`
	Members  []*Member `json:"members" db:"members" binding:"required,dive"`
}

type Member struct {
	UserID   uint   `json:"user_id" db:"user_id" binding:"required"`
	UserName string `json:"username" db:"username" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active" binding:"required"`
}
