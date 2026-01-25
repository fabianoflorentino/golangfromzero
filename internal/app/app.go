package app

import (
	"context"
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/internal/webserver"
)

// Run initializes and starts the application.
func Run() error {
	ctx := context.Background()
	logger := NewLogger()

	pool, err := NewDatabasePool(ctx, logger)
	if err != nil {
		slog.ErrorContext(ctx, "database init failed", "error", err)
		return err
	}
	defer pool.Close()

	slog.InfoContext(ctx, "database connection pool initialized")

	c := NewContainer(pool, logger)

	if err := webserver.Start(ctx, c.Router); err != nil {
		slog.ErrorContext(ctx, "http server start failed", "error", err)
		return err
	}

	slog.InfoContext(ctx, "server exited gracefully")

	return nil
}
