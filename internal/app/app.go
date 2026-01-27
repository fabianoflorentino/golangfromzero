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

	go func() {
		if err := webserver.Start(ctx, c.Router); err != nil {
			logger.ErrorContext(ctx, "http server start failed", "error", err)
		}
	}()
	<-ctx.Done()

	if err := c.Shutdown(ctx); err != nil {
		c.Logger.Warn("error during application shutdown", "error", err)
	}

	slog.InfoContext(ctx, "server exited gracefully")

	return nil
}
