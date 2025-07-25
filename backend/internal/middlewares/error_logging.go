package middlewares

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

func ErrorLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()

		c.Set(RequestIDKey, requestID)

		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				slog.Error("Request processing error",
					slog.String(RequestIDKey, requestID),
					slog.String("method", c.Request.Method),
					slog.String("path", c.Request.URL.Path),
					slog.Int("status_code", c.Writer.Status()),
					slog.Any("error_details", err.Err),
					slog.Int64("error_type", int64(err.Type)),
					slog.String("client_ip", c.ClientIP()),
					slog.String("user_agent", c.Request.UserAgent()),
				)
			}
		}

		if len(c.Errors) == 0 {
			duration := time.Since(start)
			slog.Info("Request completed successfully",
				slog.String(RequestIDKey, requestID),
				slog.String("method", c.Request.Method),
				slog.String("path", c.Request.URL.Path),
				slog.Int("status_code", c.Writer.Status()),
				slog.Duration("duration_ms", duration),
			)
		}
	}
}
