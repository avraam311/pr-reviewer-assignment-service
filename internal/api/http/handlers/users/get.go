package users

import (
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/responses"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetReviews(c *gin.Context) {
	userID := c.Param("user_id")

	reviews, err := h.service.GetReviews(c.Request.Context(), userID)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("failed to get reviews")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseOK(c, reviews)
}
