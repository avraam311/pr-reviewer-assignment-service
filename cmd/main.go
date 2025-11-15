package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"syscall"
	"time"

	handlerTeams "github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/teams"
	handlerUsers "github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/users"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/server"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/config"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	repositoryTeams "github.com/avraam311/pr-reviewer-assignment-service/internal/repository/teams"
	repositoryUsers "github.com/avraam311/pr-reviewer-assignment-service/internal/repository/users"
	serviceTeams "github.com/avraam311/pr-reviewer-assignment-service/internal/usecase/teams"
	serviceUsers "github.com/avraam311/pr-reviewer-assignment-service/internal/usecase/users"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	configFilePath = "config/local.yaml"
	envFilePath    = ".env"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger.Init()
	cfg := config.New()
	if err := cfg.LoadEnvFiles(envFilePath); err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to load env file")
	}
	cfg.EnableEnv("")
	if err := cfg.LoadConfigFiles(configFilePath); err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to load config file")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.GetString("DB_HOST"), cfg.GetInt("DB_PORT"), cfg.GetString("DB_USER"),
		cfg.GetString("DB_PASSWORD"), cfg.GetString("DB_NAME"), cfg.GetString("DB_SSL_MODE"),
	)
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to parse pgxpool dsn")
	}
	poolConfig.MaxConns = cfg.GetInt32("db.max_conns")
	poolConfig.MaxConnLifetime = cfg.GetDuration("db.max_conn_lifetime")
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to init pgxpool")
	}
	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		logger.Logger.Fatal().Err(err).Msg("failed to init pgxpool")
	}

	repoTeams := repositoryTeams.New(pool)
	srvcTeams := serviceTeams.New(repoTeams)
	handTeams := handlerTeams.New(srvcTeams)
	repoUsers := repositoryUsers.New(pool)
	srvcUsers := serviceUsers.New(repoUsers)
	handUsers := handlerUsers.New(srvcUsers)

	router := server.NewRouter(cfg, handTeams, handUsers)
	srv := server.NewServer(cfg.GetString("server.port"), router)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Logger.Fatal().Err(err).Msg("failed to run server")
		}
	}()
	logger.Logger.Info().Msg("server is running")

	<-ctx.Done()
	logger.Logger.Info().Msg("shutdown signal received")

	shutdownCtx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	logger.Logger.Info().Msg("shutting down")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to shutdown server")
	}
	if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
		logger.Logger.Info().Msg("timeout exceeded, forcing shutdown")
	}

	repoTeams.Close()
}
