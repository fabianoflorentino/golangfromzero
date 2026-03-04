package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDatabasePool initializes and returns a new database connection pool.
func NewDatabasePool(ctx context.Context, logger *slog.Logger) (*pgxpool.Pool, error) {
	databaseConfig, err := database.LoadConnectionPoolConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}

	db, err := database.NewConnectionPool(ctx, databaseConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create database pool: %w", err)
	}

	logger.InfoContext(ctx, "database connection established")

	return db, nil
}
