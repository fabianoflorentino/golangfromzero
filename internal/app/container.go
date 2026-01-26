package app

import (
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/internal/controllers"
	"github.com/fabianoflorentino/golangfromzero/internal/webserver"
	"github.com/fabianoflorentino/golangfromzero/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	Router *webserver.Router
}

func NewContainer(db *pgxpool.Pool, log *slog.Logger) *Container {
	userRepo := repository.NewUserRepository(db, log)
	userController := controllers.NewUserController(userRepo, log)

	healthController := controllers.NewHealthController(log)

	handlers := webserver.Handlers{
		User:   userController,
		Health: healthController,
	}

	router := webserver.NewRouter(handlers, log)

	return &Container{
		Router: router,
	}
}
