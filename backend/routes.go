package main

import (
	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/handlers"
	"net/http"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Golang ;)"))
	})
	r.Get("/api/v1/register", handlers.RegisterNewUser)

	r.Route("/api/v1/todos/", func(r chi.Router) {
		r.Get("/", handlers.GetTodos)
		r.Post("/", handlers.CreateTodo)
		r.Route("/:todoId", func(r chi.Router) {
			r.Get("/", handlers.GetTodo)
			r.Put("/", handlers.PutTodo)
			r.Patch("/", handlers.PatchTodo)
			r.Delete("/", handlers.DeleteTodo)
		})

	})

	r.Route("/api/v1/users/", func(r chi.Router) {
		r.Get("/", handlers.GetUsers)
		r.Post("/", handlers.CreateUser)
		r.Route("/:userId", func(r chi.Router) {
			r.Get("/", handlers.GetUser)
			r.Put("/", handlers.PutUser)
			r.Patch("/", handlers.PatchUser)
			r.Delete("/", handlers.DeleteUser)
		})
	})

	return r
}
