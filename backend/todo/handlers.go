package todo

import (
	"net/http"
	"encoding/json"
	"github.com/pressly/chi"
)

func ListTodos(w http.ResponseWriter, r *http.Request) {
	todos := RepoListTodo()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	todo := RepoFindTodo(todoId)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todo); err != nil {
		panic(err)
	}
}

func CreateTodo(w http.ResponseWriter, req *http.Request) {

}

func PutTodo(w http.ResponseWriter, req *http.Request) {

}

func PatchTodo(w http.ResponseWriter, req *http.Request) {

}

func DeleteTodo(w http.ResponseWriter, req *http.Request) {

}
