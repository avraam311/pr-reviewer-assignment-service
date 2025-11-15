package teams

import (
	"errors"
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/responses"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/repository/teams"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTeam(c *gin.Context) {
	teamName := c.Param("team_name")

	team, err := h.service.GetTeam(c.Request.Context(), teamName)
	if err != nil {
		if errors.Is(err, teams.ErrTeamNotFound) {
			logger.Logger.Error().Err(err).Interface("team_name", teamName).Msg("team.team_name not found")
			responses.ResponseError(c, responses.ErrCodeNotFound, "team.team_name not found", http.StatusBadRequest)
			return
		}

		logger.Logger.Error().Err(err).Msg("failed to get team")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseOK(c, team)
}
