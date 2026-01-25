package webserver

import (
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/internal/middleware"
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

	r.Use(middleware.LoggingMiddleware(log))

	RegisterUserRoutes(r, handlers.User)

	return &Router{Router: r}
}
