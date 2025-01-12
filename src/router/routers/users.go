package routers

import (
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/controllers"
)

const userEndpoint = "/users"

var userRouters = []Routers{
	{
		URI:              userEndpoint,
		Method:           http.MethodPost,
		Function:         controllers.CreateUser,
		VerificationAuth: false,
	},
	{
		URI:              userEndpoint + "/{id}",
		Method:           http.MethodGet,
		Function:         controllers.FindUser,
		VerificationAuth: false,
	},
	{
		URI:              userEndpoint,
		Method:           http.MethodGet,
		Function:         controllers.FindUsers,
		VerificationAuth: false,
	},
	{
		URI:              userEndpoint,
		Method:           http.MethodPut,
		Function:         controllers.UpdateUser,
		VerificationAuth: false,
	},
	{
		URI:              userEndpoint,
		Method:           http.MethodDelete,
		Function:         controllers.DeleteUser,
		VerificationAuth: false,
	},
}
