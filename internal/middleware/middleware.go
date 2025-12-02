package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		if statusCode >= 400 {
			log.Warn().
				Str("method", method).
				Str("path", path).
				Int("status", statusCode).
				Dur("duration", duration).
				Str("client_ip", c.ClientIP()).
				Msg("Request completed with error")
		} else {
			log.Info().
				Str("method", method).
				Str("path", path).
				Int("status", statusCode).
				Dur("duration", duration).
				Msg("Request completed")
		}
	}
}
