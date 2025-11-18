package server

import (
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/pr"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/statistics"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/teams"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers/users"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/config"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg *config.Config, handlerTeam *teams.Handler, handlerUser *users.Handler, handlerPR *pr.Handler, handlerStats *statistics.Handler) *gin.Engine {
	e := gin.Default()

	teamGroup := e.Group("/team")
	{
		teamGroup.POST("/add", handlerTeam.AddTeam)
		teamGroup.GET("/get/:team_name", handlerTeam.GetTeam)
		teamGroup.POST("/deactivateTeamUsers/:team_name", handlerTeam.DeactivateTeamUsers)
	}
	usersGroup := e.Group("/users")
	{
		usersGroup.POST("/setIsActive", handlerUser.SetIsActive)
		usersGroup.GET("/getReview/:user_id", handlerUser.GetReviews)
	}
	prGroup := e.Group("/pullRequest")
	{
		prGroup.POST("/create", handlerPR.CreatePR)
		prGroup.POST("/merge", handlerPR.MergePR)
		prGroup.POST("/reassign", handlerPR.ReassignPR)
	}
	statsGroup := e.Group("/statistics")
	{
		statsGroup.GET("/getPRsForUser/:user_id", handlerStats.GetStatistics)
	}

	return e
}

func NewServer(addr string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
