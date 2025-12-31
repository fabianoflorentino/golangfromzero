package router

import (
	"github.com/fabianoflorentino/golangfromzero/src/routes"
	"github.com/gorilla/mux"
)

// NewRouter creates and returns a new Gorilla Mux router instance.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
