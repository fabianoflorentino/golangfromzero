package routes

import (
	"net/http"

	"github.com/fabianoflorentino/golangfromzero/src/controllers"
)

func UserRouters(userController *controllers.UserController) []Route {

	return []Route{
		{
			URI:           "/users",
			Method:        http.MethodPost,
			Function:      userController.Create,
			Authenticated: false,
		},
		{
			URI:           "/users",
			Method:        http.MethodGet,
			Function:      userController.SearchByName,
			Authenticated: false,
		},
		{
			URI:           "/users/{userID}",
			Method:        http.MethodGet,
			Function:      userController.SearchByID,
			Authenticated: false,
		},
		{
			URI:           "/users/{userID}",
			Method:        http.MethodPut,
			Function:      userController.Update,
			Authenticated: false,
		},
		{
			URI:           "/users/{userID}",
			Method:        http.MethodDelete,
			Function:      userController.Delete,
			Authenticated: false,
		},
	}
}
