package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// getLogger assigns the logger based on configuration
func getLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s\n%s",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.ErrorMessage,
		)

		// param.Request.Proto
		// param.Request.UserAgent(),
	})
}
