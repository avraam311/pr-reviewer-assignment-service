package statistics

import (
	"errors"
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/responses"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/repository/statistics"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetStatistics(c *gin.Context) {
	userID := c.Param("user_id")

	stats, err := h.service.GetStatistics(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, statistics.ErrUserNotFound) {
			logger.Logger.Error().Err(err).Interface("user_id", userID).Msg("user not found")
			responses.ResponseError(c, responses.ErrCodeNotFound, "user not found", http.StatusBadRequest)
			return
		}

		logger.Logger.Error().Err(err).Msg("failed to get statistics")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseOK(c, stats)
}
