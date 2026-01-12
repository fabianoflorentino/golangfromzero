package router

import (
	"github.com/fabianoflorentino/golangfromzero/src/controllers"
	"github.com/fabianoflorentino/golangfromzero/src/helper"
	"github.com/fabianoflorentino/golangfromzero/src/routes"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewRouter creates and returns a new Gorilla Mux router instance.
func NewRouter(cfg helper.Config, db *pgxpool.Pool) *mux.Router {
	userController := controllers.NewUserController(cfg, db)

	r := mux.NewRouter()
	return routes.Configure(r, userController)
}
