package webserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, h UserHandler) {
	users := r.PathPrefix("/users").Subrouter()

	users.HandleFunc("", h.Create).Methods(http.MethodPost)
	users.HandleFunc("", h.SearchByName).Methods(http.MethodGet)

	users.HandleFunc("/{userID}", h.SearchByID).Methods(http.MethodGet)
	users.HandleFunc("/{userID}", h.Update).Methods(http.MethodPut)
	users.HandleFunc("/{userID}", h.Delete).Methods(http.MethodDelete)
}
