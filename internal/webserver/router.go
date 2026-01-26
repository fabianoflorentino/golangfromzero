package webserver

import (
	"log/slog"

	"github.com/fabianoflorentino/golangfromzero/internal/middleware"
	"github.com/gorilla/mux"
)

type Handlers struct {
	User   UserHandler
	Health HealthHandler
}

type Router struct {
	*mux.Router
}

func NewRouter(handlers Handlers, log *slog.Logger) *Router {
	r := mux.NewRouter()

	r.Use(middleware.LoggingMiddleware(log))

	RegisterRoutesUser(r, handlers.User)
	RegisterRoutesHealth(r, handlers.Health)

	return &Router{Router: r}
}
