package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/database"
	"github.com/tokillamockingbird/golang-todo/backend/headers"
	"github.com/tokillamockingbird/golang-todo/backend/models"
	e "github.com/tokillamockingbird/golang-todo/backend/errors"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	CreateUser(w, r)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(database.GetUsers()); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentType(w, http.StatusOK)
	userId := chi.URLParam(r, "userId")
	if err := json.NewEncoder(w).Encode(database.GetUser(userId)); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body == nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user = database.CreateUser(user)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body == nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user = database.UpdateUser(user)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body == nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, 400)
	}
	userId := chi.URLParam(r, "userId")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, 400)
		return
	}

	user = database.PatchUser(userId, user)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	database.DeleteUser(userId)

	headers.SetJSONContentType(w, http.StatusOK)
}
