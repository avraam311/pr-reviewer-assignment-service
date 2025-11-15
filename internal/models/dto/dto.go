package dto

type Team struct {
	TeamName string    `json:"team_name" db:"team_name" validate:"required"`
	Members  []*Member `json:"members" db:"members" validate:"required"`
}

type Member struct {
	UserID   uint   `json:"user_id" db:"user_id" validate:"requred"`
	UserName string `json:"username" db:"username" validate:"requred"`
	IsActive bool   `json:"is_active" db:"is_active" validate:"required"`
}
