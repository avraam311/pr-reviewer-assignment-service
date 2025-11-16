package db

import "time"

type User struct {
	UserID   string `json:"user_id" db:"user_id" binding:"required"`
	UserName string `json:"username" db:"username" binding:"required"`
	TeamName string `json:"team_name" db:"team_name" binding:"required"`
	IsActive bool   `json:"is_active" db:"is_active" binding:"required"`
}

type PR struct {
	PRID              string   `json:"pull_request_id" db:"pull_request_id" binding:"required"`
	PRName            string   `json:"pull_request_name" db:"pull_request_name" binding:"required"`
	AuthorID          string   `json:"author_id" db:"author_id" binding:"required"`
	Status            string   `json:"status" db:"status" binding:"required"`
	AssignedReviewers []string `json:"assigned_reviewers" db:"assigned_reviewers" binding:"required"`
}

type PRWithMergedAt struct {
	PRID              string    `json:"pull_request_id" db:"pull_request_id" binding:"required"`
	PRName            string    `json:"pull_request_name" db:"pull_request_name" binding:"required"`
	AuthorID          string    `json:"author_id" db:"author_id" binding:"required"`
	Status            string    `json:"status" db:"status" binding:"required"`
	AssignedReviewers []string  `json:"assigned_reviewers" db:"assigned_reviewers" binding:"required"`
	MergedAt          time.Time `json:"merged_at" db:"merged_at" binding:"required"`
}

type PRWithReplacedBy struct {
	PR         *PR    `json:"pr" db:"pr" binding:"required"`
	ReplacedBy string `json:"replaced_by" db:"replaced_by" binding:"required"`
}

type PRShort struct {
	PRID     string `json:"pull_request_id" db:"pull_request_id" binding:"required"`
	PRName   string `json:"pull_request_name" db:"pull_request_name" binding:"required"`
	AuthorID string `json:"author_id" db:"author_id" binding:"required"`
	Status   string `json:"status" db:"status" binding:"required"`
}
