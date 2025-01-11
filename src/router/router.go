package router

import (
	"github.com/gorilla/mux"

	"github.com/felipecesargomes/social-apiGO/src/router/routers"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.ConfigRouters(r)
}
