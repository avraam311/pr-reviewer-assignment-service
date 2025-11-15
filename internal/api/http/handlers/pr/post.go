package pr

import (
	// "errors"
	"fmt"
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/api/http/responses"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"
	"github.com/avraam311/pr-reviewer-assignment-service/internal/models/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePR(c *gin.Context) {
	var pr dto.PR
	if err := c.ShouldBindJSON(&pr); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to decode or validate request body")
		responses.ResponseError(c, responses.ErrInvalidJSON, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	prResp, err := h.service.CreatePR(c.Request.Context(), &pr)
	if err != nil {
		// if errors.Is(err, pr.ErrAuthorNotFound) {
		// 	logger.Logger.Error().Err(err).Interface("pr", pr).Msg("pr.author_id not found")
		// 	responses.ResponseError(c, responses.ErrCodeNotFound, "pr.author_id not found", http.StatusBadRequest)
		// 	return
		// } else if errors.Is(err, pr.ErrPRAlreadyExists) {
		// 	logger.Logger.Error().Err(err).Interface("pr", pr).Msg("pr.pull_request_id or pr.pull_request_name already exists")
		// 	responses.ResponseError(c, responses.ErrCodePrExists, "pr.pull_request_id or pr.pull_request_name already exists", http.StatusConflict)
		// 	return
		// }

		logger.Logger.Error().Err(err).Interface("pr", pr).Msg("failed to create pr")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseCreated(c, prResp)
}

func (h *Handler) MergePR(c *gin.Context) {
	var pr dto.PRWithPRID
	if err := c.ShouldBindJSON(&pr); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to decode or validate request body")
		responses.ResponseError(c, responses.ErrInvalidJSON, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	prResp, err := h.service.MergePR(c.Request.Context(), &pr)
	if err != nil {
		// if errors.Is(err, pr.ErrPRNotFound) {
		// 	logger.Logger.Error().Err(err).Interface("pr", pr).Msg("pr.pull_request_id not found")
		// 	responses.ResponseError(c, responses.ErrCodeNotFound, "pr.pull_request_id not found", http.StatusBadRequest)
		// 	return
		// }

		logger.Logger.Error().Err(err).Interface("pr", pr).Msg("failed to merge pr")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseOK(c, prResp)
}

func (h *Handler) ReassignPR(c *gin.Context) {
	var pr dto.PRWithOldUserID
	if err := c.ShouldBindJSON(&pr); err != nil {
		logger.Logger.Error().Err(err).Msg("failed to decode or validate request body")
		responses.ResponseError(c, responses.ErrInvalidJSON, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	prResp, err := h.service.ReassignPR(c.Request.Context(), &pr)
	if err != nil {
		// if errors.Is(err, pr.ErrPROrOldUserNotFound) {
		// 	logger.Logger.Error().Err(err).Interface("pr", pr).Msg("pr.pull_request_id or pr.old_user_id not found")
		// 	responses.ResponseError(c, responses.ErrCodeNotFound, "pr.pull_request_id or pr.old_user_id not found", http.StatusBadRequest)
		// 	return
		// } else if errors.Is(err, pr.ErrReassignAfterMerge) {
		// 	logger.Logger.Error().Err(err).Interface("pr", pr).Msg("pr already merged")
		// 	responses.ResponseError(c, responses.ErrCodePrMerged, "pr already merged", http.StatusConflict)
		// 	return

		logger.Logger.Error().Err(err).Interface("pr", pr).Msg("failed to reassign pr")
		responses.ResponseError(c, responses.ErrInternalServer, "internal server error", http.StatusInternalServerError)
		return
	}

	responses.ResponseOK(c, prResp)
}
