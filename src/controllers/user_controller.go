package controllers

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user"))
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
