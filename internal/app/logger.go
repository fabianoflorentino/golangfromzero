package app

import (
	"log/slog"
	"os"
)

// NewLogger creates and returns a new structured logger.
func NewLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return logger
}
