package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/config"

	"github.com/felipecesargomes/social-apiGO/src/router"
)

func main() {
	env := config.Load()
	r := router.Generate()
	msgTitle := fmt.Sprintf("API running on port %s", env.API_PORT)
	fmt.Println(msgTitle)
	fmt.Println(env.DB_TYPE, env.DB_HOST, env.DB_USER, env.DB_PASS, env.DB_NAME, env.DB_PORT)
	_, err := config.Connect(env.DB_TYPE, env.DB_HOST, env.DB_USER, env.DB_PASS, env.DB_NAME, env.DB_PORT)
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	fmt.Println("Connected to database")
	log.Fatal(http.ListenAndServe(":5000", r))

}
