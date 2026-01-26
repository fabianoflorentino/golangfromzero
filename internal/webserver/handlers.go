package webserver

import "net/http"

type UserHandler interface {
	Create(http.ResponseWriter, *http.Request)
	SearchByName(http.ResponseWriter, *http.Request)
	SearchByID(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

type HealthHandler interface {
	Check(http.ResponseWriter, *http.Request)
}
