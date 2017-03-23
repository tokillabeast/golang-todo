package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/database"
	e "github.com/tokillamockingbird/golang-todo/backend/errors"
	"github.com/tokillamockingbird/golang-todo/backend/headers"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

var TodosHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { // FIXME: migrate to handlers???
	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(database.GetTodos()); err != nil {
		http.Error(w, err.Error(), 400) // FIXME: new function in errors like CheckAndLogError
		return
	}
})

func GetTodo(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentType(w, http.StatusOK)
	todoId := chi.URLParam(r, "todoId")
	if err := json.NewEncoder(w).Encode(database.GetTodo(todoId)); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if r.Body == nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	todo = database.CreateTodo(todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func PutTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if r.Body == nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, 400)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	todo = database.UpdateTodo(todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func PatchTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if r.Body == nil {
		http.Error(w, e.EmptyRequestBodyErrorMsg, 400)
		return
	}
	todoId := chi.URLParam(r, "todoId") // FIXME: check that todoId contain value
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	todo = database.PatchTodo(todoId, todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	database.DeleteTodo(todoId)

	headers.SetJSONContentType(w, http.StatusOK)
}
