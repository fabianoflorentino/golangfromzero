package router

import (
	"github.com/fabianoflorentino/golangfromzero/src/controllers"
	"github.com/fabianoflorentino/golangfromzero/src/helper"
	"github.com/fabianoflorentino/golangfromzero/src/routes"
	"github.com/gorilla/mux"
)

// NewRouter creates and returns a new Gorilla Mux router instance.
func NewRouter() *mux.Router {
	cfg := helper.ConfigTimeout
	userController := controllers.NewUserController(cfg)

	r := mux.NewRouter()
	return routes.Configure(r, userController)
}
