package app

import (
	"context"
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/internal/controllers"
	"github.com/fabianoflorentino/golangfromzero/internal/webserver"
	"github.com/fabianoflorentino/golangfromzero/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	Router *webserver.Router
	DB     *pgxpool.Pool
	Logger *slog.Logger
}

func NewContainer(db *pgxpool.Pool, log *slog.Logger) *Container {
	handlers := webserver.Handlers{
		User:   buildUserController(db, log),
		Health: buildHealthController(log),
	}

	return &Container{
		Router: webserver.NewRouter(handlers, log),
		DB:     db,
		Logger: log,
	}
}

func (c *Container) Shutdown(ctx context.Context) error {
	c.Logger.Info("Shutting down application container")

	if c.DB != nil {
		c.Logger.Info("Closing database connection pool")
		c.DB.Close()
	}

	return nil
}

func buildUserController(db *pgxpool.Pool, log *slog.Logger) *controllers.UserController {
	userRepo := repository.NewUserRepository(db, log)
	return controllers.NewUserController(userRepo, log)
}

func buildHealthController(log *slog.Logger) *controllers.HealthController {
	return controllers.NewHealthController(log)
}
