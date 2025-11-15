package teams

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/handlers"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddTeam(c *gin.Context) {
	var team dto.Team
	if err := json.NewDecoder(c.Request.Body).Decode(&team); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to decode request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	if err := h.validator.Struct(team); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to validate request body")
		handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("validation error: %s", err.Error()))
		return
	}

	err := h.service.AddTeam(c.Request.Context(), &team)
	if err != nil {
		// if errors.Is(err, teams.ErrTeamAlreadyExists) {
		// 	logger.Logger.Error().Err(err).Interface("team", team).Msg("team_name already exists")
		// 	handlers.Fail(c.Writer, http.StatusBadRequest, fmt.Errorf("team_name already exists"))
		// 	return
		// }

		logger.Logger.Error().Err(err).Interface("team", team).Msg("failed to add team")
		handlers.Fail(c.Writer, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	handlers.Created(c.Writer, team)
}
