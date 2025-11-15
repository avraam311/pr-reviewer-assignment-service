package teams

import (
	"context"
	"errors"
	"fmt"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrTeamAlreadyExists = errors.New("team already exists")
)

type Repository struct {
	db  *pgxpool.Pool
	cfg *config.Config
}

func New(cfg *config.Config) (*Repository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetString("DB_HOST"), cfg.GetInt("DB_PORT"), cfg.GetString("DB_user"),
		cfg.GetString("DB_PASSWORD"), cfg.GetString("DB_NAME"), cfg.GetString("DB_SSL"),
	)
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgxpool config: %w", err)
	}

	poolConfig.MaxConns = cfg.GetInt32("db.max_conns")
	poolConfig.MaxConnLifetime = cfg.GetDuration("db.max_conn_lifetime")
	poolConfig.MaxConnIdleTime = cfg.GetDuration("db.max_conn_idletime")

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &Repository{
		db:  pool,
		cfg: cfg}, nil
}

func (r *Repository) Close() {
	r.db.Close()
}
