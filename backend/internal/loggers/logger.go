package loggers

import (
	"log/slog"
	"os"
)

func Init() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
		// AddSource: true,
	})
	slog.SetDefault(slog.New(handler))
}
