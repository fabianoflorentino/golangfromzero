package app

import (
	"context"
	"log"
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDatabasePool initializes and returns a new database connection pool.
func NewDatabasePool(ctx context.Context, logger *slog.Logger) (*pgxpool.Pool, error) {
	databaseConfig, err := database.LoadConnectionPoolConfig()
	if err != nil {
		log.Fatalf("failed to load database config: %s", err)
	}

	db, err := database.NewConnectionPool(ctx, databaseConfig)
	if err != nil {
		log.Fatalf("failed to create database pool: %s", err)
	}

	logger.InfoContext(ctx, "database connection established")

	return db, nil
}
