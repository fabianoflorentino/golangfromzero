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
		Function:      controllers.SearchByName,
		Authenticated: false,
	},
	{
		URI:           "/users/{userID}",
		Method:        http.MethodGet,
		Function:      controllers.SearchByID,
		Authenticated: false,
	},
	{
		URI:           "/users/{userID}",
		Method:        http.MethodPut,
		Function:      controllers.Update,
		Authenticated: false,
	},
	{
		URI:           "/users/{userID}",
		Method:        http.MethodDelete,
		Function:      controllers.Delete,
		Authenticated: false,
	},
}
