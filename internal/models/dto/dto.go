package dto

type Team struct {
	TeamName string  `json:"team_name" db:"team_name" binding:"required"`
	Members  []*User `json:"members" db:"members" binding:"required,dive"`
}

type User struct {
	UserID   string `json:"user_id" db:"user_id" binding:"required"`
	UserName string `json:"username" db:"username" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active" binding:"required"`
}

type UserWithIsActive struct {
	UserID   string `json:"user_id" db:"user_id" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active" binding:"required"`
}
