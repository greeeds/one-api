package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/greeeds/one-api/common/helper"
	"github.com/greeeds/one-api/common/logger"
)

func abortWithMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"error": gin.H{
			"message": helper.MessageWithRequestId(message, c.GetString(logger.RequestIdKey)),
			"type":    "one_api_error",
		},
	})
	c.Abort()
	logger.Error(c.Request.Context(), message)
}
