package middlewares

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

func ErrorLoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()

		ctx.Set(RequestIDKey, requestID)

		ctx.Next()

		if len(ctx.Errors) > 0 {
			for _, err := range ctx.Errors {
				slog.Error("Request processing error",
					slog.String(RequestIDKey, requestID),
					slog.String("method", ctx.Request.Method),
					slog.String("path", ctx.Request.URL.Path),
					slog.Int("status_code", ctx.Writer.Status()),
					slog.Any("error_details", err.Err),
					slog.Int64("error_type", int64(err.Type)),
					slog.String("client_ip", ctx.ClientIP()),
					slog.String("user_agent", ctx.Request.UserAgent()),
				)
			}
		}

		if len(ctx.Errors) == 0 {
			duration := time.Since(start)
			slog.Info("Request completed successfully",
				slog.String(RequestIDKey, requestID),
				slog.String("method", ctx.Request.Method),
				slog.String("path", ctx.Request.URL.Path),
				slog.Int("status_code", ctx.Writer.Status()),
				slog.Duration("duration_ms", duration),
			)
		}
	}
}
