package main

import (
	"context"
	"errors"
	"os/signal"
	"syscall"
	"time"

	handlerTeams "github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/teams"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/server"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/config"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	repositoryTeams "github.com/avraam311/pr-reviewer-assignment-service/internal/repository/teams"
	serviceTeams "github.com/avraam311/pr-reviewer-assignment-service/internal/usecase/teams"
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

	repoTeams, err := repositoryTeams.New(cfg)
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("failed to init teams repo")
	}
	srvcTeams := serviceTeams.New(repoTeams)
	handTeams := handlerTeams.New(srvcTeams)

	router := server.NewRouter(cfg, handTeams)
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
