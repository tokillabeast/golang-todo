package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tokillamockingbird/golang-todo/backend/database"
	"github.com/tokillamockingbird/golang-todo/backend/headers"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

func RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Body != nil {
		http.Error(w, "Please send a request body", 400)
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
	}
}
