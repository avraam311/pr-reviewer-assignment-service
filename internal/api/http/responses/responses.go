package responses

import (
	"net/http"

	"github.com/avraam311/pr-reviewer-assignment-service/internal/infra/logger"

	"github.com/gin-gonic/gin"
)

const (
	ErrCodeTeamExists  = "TEAM_EXISTS"
	ErrCodePrExists    = "PR_EXISTS"
	ErrCodePrMerged    = "PR_MERGED"
	ErrCodeNotAssigned = "NOT_ASSIGNED"
	ErrCodeNoCandidate = "NO_CANDIDATE"
	ErrCodeNotFound    = "NOT_FOUND"
)

type Success struct {
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func ResponseOK(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, Success{Result: result})
}

func ResponseCreated(c *gin.Context, result interface{}) {
	c.JSON(http.StatusCreated, Success{Result: result})
}

func ResponseError(c *gin.Context, code string, message string, statusCode int) {
	resp := ErrorResponse{}
	resp.Error.Code = code
	resp.Error.Message = message
	c.JSON(statusCode, resp)
	c.Abort()
}

func HandleError(c *gin.Context, err error, code string, message string, statusCode int) {
	logger.Logger.Error().Err(err).Msg(message)
	ResponseError(c, code, message, statusCode)
}
