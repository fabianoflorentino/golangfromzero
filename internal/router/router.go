package router

import (
	"github.com/fabianoflorentino/golangfromzero/internal/controllers"
	"github.com/fabianoflorentino/golangfromzero/internal/routes"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewRouter creates and returns a new Gorilla Mux router instance.
func NewRouter(db *pgxpool.Pool) *mux.Router {
	userController := controllers.NewUserController(db)

	r := mux.NewRouter()
	return routes.Configure(r, userController)
}
