package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/config"
	"github.com/felipecesargomes/social-apiGO/src/models"
	"github.com/felipecesargomes/social-apiGO/src/repositories"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err := json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db := config.Connection

	repository := repositories.NewRepositoryUser(db)
	userID, err := repository.CreateUser(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Id insert: %d", userID)))

}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find user"))
}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find all users"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user"))
}
