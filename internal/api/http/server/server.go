package server

import (
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/config"

	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, handlerTeam *teams.Handler) *gin.Engine {
	e := gin.Default()

	api := e.Group("/api/v1")
	teamsGroup := api.Group("/team")
	{
		teamsGroup.POST("/add", handlerTeam.Add)
		teamsGroup.GET("/get", handlerTeam.Get)
	}

	return e
}

func NewServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
