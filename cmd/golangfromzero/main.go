package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/fabianoflorentino/golangfromzero/internal/server"
	"github.com/fabianoflorentino/golangfromzero/src/helper"
	"github.com/fabianoflorentino/golangfromzero/src/router"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error load environment variables:  %v", err.Error())
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := helper.ConfigTimeout
	databaseConfig := database.LoadConnectionPoolConfig()

	ctx := context.Background()

	pool, err := database.NewConnectionPool(ctx, databaseConfig)
	if err != nil {
		log.Fatalf("falied to create database pool: %s", err)
	}
	defer pool.Close()

	slog.Info("database connection pool initialized successfully")

	r := router.NewRouter(cfg, pool)

	if err := server.Start(ctx, r); err != nil {
		log.Fatalf("failed to start http server: %s", err)
	}

	pool.Close()
	slog.Info("server exited gracefully")
}
