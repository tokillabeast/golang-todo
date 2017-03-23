package routes

import (
	"net/http"

	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/handlers"
	"github.com/tokillamockingbird/golang-todo/backend/middleware"
)

type TodoResource struct{}

func (t TodoResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", middleware.JwtMiddleware.Handler(handlers.TodosHandler).(http.HandlerFunc))
	r.Post("/", handlers.CreateTodo)
	r.Route("/:todoId", func(r chi.Router) {
		r.Get("/", handlers.GetTodo)
		r.Put("/", handlers.PutTodo)
		r.Patch("/", handlers.PatchTodo)
		r.Delete("/", handlers.DeleteTodo)
	})

	return r
}
