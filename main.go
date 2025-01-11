package main

import (
	"log"
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/router"
)

func main() {
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5080", r))
}
