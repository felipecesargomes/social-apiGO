package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/felipecesargomes/social-apiGO/src/config"
	"github.com/felipecesargomes/social-apiGO/src/models"
	"github.com/felipecesargomes/social-apiGO/src/repositories"
	"github.com/felipecesargomes/social-apiGO/src/responses"
	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read body request
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

	if err := user.Prepare("add"); err != nil {
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

func FindUsers(w http.ResponseWriter, r *http.Request) {
	userParam := r.URL.Query().Get("user")

	db := config.Connection
	repository := repositories.NewRepositoryUser(db)

	users, err := repository.Find(userParam)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	paramenters := mux.Vars(r)
	userID, err := strconv.ParseUint(paramenters["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	if err := user.Prepare("update"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db := config.Connection

	repository := repositories.NewRepositoryUser(db)
	if err := repository.UpdateUser(userID, user); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db := config.Connection

	repository := repositories.NewRepositoryUser(db)
	if err := repository.DeleteUser(userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db := config.Connection
	defer db.Close()

	repository := repositories.NewRepositoryUser(db)
	user, err := repository.FindById(userID)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
