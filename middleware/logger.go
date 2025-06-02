package middleware

import (
	"api/train/infra/logger"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.URL.Path

		msg := fmt.Sprintf("[%s] %s %s | %d | %s | %s",
			start.Format(time.RFC3339),
			method,
			path,
			status,
			clientIP,
			userAgent,
		)
		logger.LogToFileAsync(msg + " | latency: " + latency.String())
	}
}
