package users

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrUserNotFound = errors.New("user not found")
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
