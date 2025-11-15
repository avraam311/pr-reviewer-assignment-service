package server

import (
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/pr"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/teams"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/users"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/config"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, handlerTeam *teams.Handler, handlerUser *users.Handler, handlerPR *pr.Handler) *gin.Engine {
	e := gin.Default()

	api := e.Group("/api/v1")
	teamGroup := api.Group("/team")
	{
		teamGroup.POST("/add", handlerTeam.AddTeam)
		teamGroup.GET("/get/:team_name", handlerTeam.GetTeam)
	}
	usersGroup := api.Group("/users")
	{
		usersGroup.POST("/setIsActive", handlerUser.SetIsActive)
	}
	prGroup := api.Group("/pullRequest")
	{
		prGroup.POST("/create", handlerPR.CreatePR)
		prGroup.POST("/merge", handlerPR.MergePR)
		prGroup.POST("/reassign", handlerPR.ReassignPR)
	}

	return e
}

func NewServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
