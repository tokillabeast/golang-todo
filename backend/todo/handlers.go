package todo

import (
	"encoding/json"
	"net/http"

	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/headers"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

func ListTodos(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(RepoListTodo()); err != nil {
		panic(err)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(RepoFindTodo(chi.URLParam(r, "todoId"))); err != nil {
		panic(err)
	}
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	todo = RepoCreateTodo(todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func PutTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&todo)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	todo = RepoUpdateTodo(todo)

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func PatchTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId") // FIXME: not working correct
	todo := RepoPatchTodo(todoId, r.Body)
	defer r.Body.Close()

	headers.SetJSONContentType(w, http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	RepoDeleteTodo(todoId)

	headers.SetJSONContentType(w, http.StatusOK)
}
