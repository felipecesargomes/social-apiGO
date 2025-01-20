package main

import (
	"log"
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/router"
	"github.com/rs/cors"
)

func main() {
	r := router.Generate()

	handler := cors.Default().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
