package teams

import (
	// "errors"
	"fmt"
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetTeam(c *gin.Context) {
	teamName := c.Param("team_name")

	team, err := h.service.GetTeam(c.Request.Context(), teamName)
	if err != nil {
		// if errors.Is(err, teams.ErrTeamNotFound) {
		// 	logger.Logger.Error().Err(err).Interface("team_name", teamName).Msg("team_name not found")
		// 	handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("team_name not found"))
		// 	return
		// }

		logger.Logger.Error().Err(err).Msg("failed to get team")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.OK(c.Writer, team)
}
