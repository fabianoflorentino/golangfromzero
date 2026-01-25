package webserver

import (
	"log/slog"

	"github.com/gorilla/mux"
)

type Handlers struct {
	User UserHandler
}

type Router struct {
	*mux.Router
}

func NewRouter(handlers Handlers, log *slog.Logger) *Router {
	r := mux.NewRouter()

	RegisterUserRoutes(r, handlers.User)

	return &Router{Router: r}
}
