package router

import (
	"fmt"

	"github.com/fabianoflorentino/golangfromzero/src/routes"
	"github.com/gorilla/mux"
)

// NewRouter creates and returns a new Gorilla Mux router instance.
func NewRouter() *mux.Router {
	fmt.Println("server up and running on port 6000")
	r := mux.NewRouter()

	return routes.Configure(r)
}
