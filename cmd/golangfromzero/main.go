package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fabianoflorentino/golangfromzero/config"
	"github.com/fabianoflorentino/golangfromzero/database"
	"github.com/fabianoflorentino/golangfromzero/src/helper"
	"github.com/fabianoflorentino/golangfromzero/src/router"
)

func main() {

	if err := config.LoadEnv(); err != nil {
		slog.Error("[ERROR]: %v", "load environment variables", err.Error())
		return
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

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		slog.Info("server starting", "address", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("sever forced to shutdown: %s", err)
	}

	pool.Close()
	slog.Info("server exited gracefully")
}
