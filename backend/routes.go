package main

import (
	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/handlers"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

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

	return r
}
