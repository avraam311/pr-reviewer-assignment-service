package teams

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/responses"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/repository/teams"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddTeam(c *gin.Context) {
	var team dto.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to decode or validate request body")
		responses.ResponseError(c, responses.ErrInvalidJSON, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	err := h.service.AddTeam(c.Request.Context(), &team)
	if err != nil {
		if errors.Is(err, teams.ErrTeamAlreadyExists) {
			logger.Logger.Error().Err(err).Interface("team", team).Msg("team.team_name already exists")
			responses.ResponseError(c, responses.ErrCodeTeamExists, "team.team_name already exists", http.StatusBadRequest)
			return
		}

		logger.Logger.Error().Err(err).Interface("team", team).Msg("failed to add team")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseCreated(c, team)
}
