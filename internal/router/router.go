package router

import (
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/internal/controllers"
	"github.com/fabianoflorentino/golangfromzero/internal/routes"
	"github.com/fabianoflorentino/golangfromzero/repository"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewRouter creates and returns a new Gorilla Mux router instance.
func NewRouter(db *pgxpool.Pool, log *slog.Logger) *mux.Router {
	userRepo := repository.NewUserRepository(db, log)
	userController := controllers.NewUserController(db, userRepo, log)

	r := mux.NewRouter()
	return routes.Configure(r, userController)
}
