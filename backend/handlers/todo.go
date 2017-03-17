package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/database"
	"github.com/tokillamockingbird/golang-todo/backend/headers"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

func ListTodos(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(database.RepoListTodo()); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentType(w, http.StatusOK)
	todoId := chi.URLParam(r, "todoId")
	if err := json.NewEncoder(w).Encode(database.RepoFindTodo(todoId)); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
	}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	todo = database.RepoCreateTodo(todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func PutTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	defer r.Body.Close()
	todo = database.RepoUpdateTodo(todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func PatchTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
	}
	todoId := chi.URLParam(r, "todoId") // FIXME: check that todoId contain value
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	todo = database.RepoPatchTodo(todoId, todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	database.RepoDeleteTodo(todoId)

	headers.SetJSONContentType(w, http.StatusOK)
}
