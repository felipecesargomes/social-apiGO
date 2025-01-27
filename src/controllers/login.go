package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/felipecesargomes/social-apiGO/src/config"
	"github.com/felipecesargomes/social-apiGO/src/models"
	"github.com/felipecesargomes/social-apiGO/src/repositories"
	"github.com/felipecesargomes/social-apiGO/src/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var userSaved models.User
	if err = json.Unmarshal(body, &userSaved); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := config.Connect("host", "port", "user", "password", "dbname", "sslmode")
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryUser(db)
	userSaved, err = repository.FindByEmail(userSaved.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(userSaved)
}
