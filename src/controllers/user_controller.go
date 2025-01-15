package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/config"
	"github.com/felipecesargomes/social-apiGO/src/models"
	"github.com/felipecesargomes/social-apiGO/src/repositories"
	"github.com/felipecesargomes/social-apiGO/src/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//Read body request
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err := json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db := config.Connection

	repository := repositories.NewRepositoryUser(db)
	userID, err := repository.CreateUser(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return

	}
	responses.JSON(w, http.StatusCreated, userID)

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
