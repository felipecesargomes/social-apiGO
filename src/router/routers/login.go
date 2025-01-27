package routers

import (
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/controllers"
)

var routerLogin = Routers{
	URI:              "/login",
	Method:           http.MethodPost,
	Function:         controllers.Login,
	VerificationAuth: false,
}
