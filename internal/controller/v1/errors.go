package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Error struct {
	Message string `json:"message"`
}

func (c *Router) Error(ctx *gin.Context, err error, msg string, statusCode int) {
	if err != nil {
		c.logger.Error("handler error", zap.Error(err))
	}

	ctx.JSON(
		statusCode,
		Error{
			Message: msg,
		},
	)
}
