package middleware

import (
	"fmt"
	"time"

	"github.com/UberionAI/coord-gin-rest/internal/logger"
	"github.com/gin-gonic/gin"
)

// RequestLogger logs HTTP requests
func RequestLogger(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Log after request processing
		duration := time.Since(start)
		statusCode := c.Writer.Status()

		msg := fmt.Sprintf("%s %s | status=%d | duration=%v",
			method, path, statusCode, duration)

		// Log based on status code
		if statusCode >= 500 {
			log.Error(msg, fmt.Errorf("server error"))
		} else if statusCode >= 400 {
			log.Warn(msg)
		} else {
			log.Info(msg)
		}
	}
}

// Recovery recovers from panics and logs them
func Recovery(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("Panic recovered", fmt.Errorf("%v", err))
				c.AbortWithStatusJSON(500, gin.H{
					"error": "Internal server error",
				})
			}
		}()
		c.Next()
	}
}

// CORS adds CORS headers (optional, for future use)
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
