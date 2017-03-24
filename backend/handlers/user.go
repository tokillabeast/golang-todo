package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/database"
	e "github.com/tokillamockingbird/golang-todo/backend/errors"
	"github.com/tokillamockingbird/golang-todo/backend/headers"
	"github.com/tokillamockingbird/golang-todo/backend/models"
	"github.com/tokillamockingbird/golang-todo/backend/services"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body == nil {
		e.Error(w, e.EmptyRequestBodyErrorMsg, http.StatusNoContent)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		e.Error(w, err.Error(), http.StatusPartialContent)
		return
	}
	user, err = database.CreateUser(user)
	if err != nil {
		e.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	token, err := services.GenerateToken(user)
	if err != nil {
		e.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(models.AuthenticateResponse{user, token}); err != nil {
		e.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var authenticate models.AuthenticateRequest
	if r.Body == nil {
		e.Error(w, e.EmptyRequestBodyErrorMsg, http.StatusNoContent)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&authenticate)
	if err != nil {
		e.Error(w, err.Error(), http.StatusPartialContent)
		return
	}
	user, err := database.GetAuthenticateUser(authenticate)
	if err != nil {
		e.Error(w, err.Error(), http.StatusNoContent)
		return
	}
	token, err := services.GenerateToken(user)
	if err != nil {
		e.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(models.AuthenticateResponse{user, token}); err != nil {
		e.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentTypeAndStatus(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(database.GetUsers()); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentTypeAndStatus(w, http.StatusOK)
	userId := chi.URLParam(r, "userId")
	if err := json.NewEncoder(w).Encode(database.GetUser(userId)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, _ = database.CreateUser(user)

	headers.SetJSONContentTypeAndStatus(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body == nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, http.StatusInternalServerError)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	user = database.UpdateUser(user)

	headers.SetJSONContentTypeAndStatus(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	headers.SetJSONContentTypeAndStatus(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	database.DeleteUser(userId)

	headers.SetJSONContentTypeAndStatus(w, http.StatusOK)
}
