package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/responses"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/repository/users"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SetIsActive(c *gin.Context) {
	var usr dto.UserWithIsActive
	if err := c.ShouldBindJSON(&usr); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to decode or validate request body")
		responses.ResponseError(c, responses.ErrInvalidJSON, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	user, err := h.service.SetIsActive(c.Request.Context(), &usr)
	if err != nil {
		if errors.Is(err, users.ErrUserNotFound) {
			logger.Logger.Error().Err(err).Interface("user", usr).Msg("user.id not found")
			responses.ResponseError(c, responses.ErrCodeNotFound, "user.id not found", http.StatusBadRequest)
			return
		}

		logger.Logger.Error().Err(err).Interface("team", usr).Msg("failed to set is_active")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseOK(c, user)
}
