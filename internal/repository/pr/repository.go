package pr

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	prStatusMerged = "MERGED"
)

var (
	ErrAuthorNotFound      = errors.New("author not found")
	ErrPRAlreadyExists     = errors.New("pr already exists")
	ErrPRNotFound          = errors.New("pr not found")
	ErrPROrOldUserNotFound = errors.New("pr or old user not found")
	ErrReassignAfterMerge  = errors.New("reassign after merge")
	ErrReviewerNotAssigned = errors.New("reviewer not assigned")
	ErrNoCandidate         = errors.New("no candidate to reassign")
)

type Repository struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		db: pool,
	}
}

func (r *Repository) Close() {
	r.db.Close()
}
