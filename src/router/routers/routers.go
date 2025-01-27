package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routers struct {
	URI              string
	Method           string
	Function         func(http.ResponseWriter, *http.Request)
	VerificationAuth bool
}

func ConfigRouters(r *mux.Router) *mux.Router {
	routers := userRouters
	routers = append(routers, routerLogin)

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}

	return r
}
