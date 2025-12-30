package routes

import "net/http"

// Route defines the structure for an API route
type Route struct {
	URI           string
	Method        string
	Function      func(http.ResponseWriter, *http.Request)
	Authenticated bool
}
