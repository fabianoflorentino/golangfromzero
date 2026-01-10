package routes

import (
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/src/controllers"
	"github.com/gorilla/mux"
)

// Route defines the structure for an API route
type Route struct {
	URI           string
	Method        string
	Function      func(http.ResponseWriter, *http.Request)
	Authenticated bool
}

// Configure sets up the routes in the provided router
func Configure(r *mux.Router, userController *controllers.UserController) *mux.Router {
	u := UserRouters(userController)

	for _, route := range u {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
