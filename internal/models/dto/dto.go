package dto

type Team struct {
	TeamName string  `json:"team_name" db:"team_name" binding:"required"`
	Members  []*User `json:"members" db:"members" binding:"required,dive"`
}

type User struct {
	UserID   string `json:"user_id" db:"user_id" binding:"required"`
	UserName string `json:"username" db:"username" binding:"required"`
	IsActive *bool  `json:"is_active" db:"is_active" binding:"required"`
}

type UserWithIsActive struct {
	UserID   string `json:"user_id" db:"user_id" binding:"required"`
	IsActive *bool  `json:"is_active" db:"is_active" binding:"required"`
}

type PR struct {
	PRID     string `json:"pull_request_id" db:"pull_request_id" binding:"required"`
	PRName   string `json:"pull_request_name" db:"pull_request_name" binding:"required"`
	AuthorID string `json:"author_id" db:"author_id" binding:"required"`
}

type PRWithPRID struct {
	PRID string `json:"pull_request_id" db:"pull_request_id" binding:"required"`
}

type PRWithOldUserID struct {
	PRID      string `json:"pull_request_id" db:"pull_request_id" binding:"required"`
	OldUserID string `json:"old_user_id" db:"old_user_id" binding:"required"`
}
