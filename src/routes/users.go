package routes

import (
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/src/controllers"
)

var UsersRoutes = []Route{
	{
		URI:           "/users",
		Method:        http.MethodPost,
		Function:      controllers.Create,
		Authenticated: false,
	},
	{
		URI:           "/users",
		Method:        http.MethodGet,
		Function:      controllers.GetAll,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodGet,
		Function:      controllers.GetByID,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodPut,
		Function:      controllers.Update,
		Authenticated: false,
	},
	{
		URI:           "/users/{userId}",
		Method:        http.MethodDelete,
		Function:      controllers.Delete,
		Authenticated: false,
	},
}
