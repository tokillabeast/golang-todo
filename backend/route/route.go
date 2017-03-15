package route

import (
	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/todo"
)

func Routes() *chi.Mux{
	r := chi.NewRouter()

	r.Route("/api/v1/todos/", func(r chi.Router){
		r.Get("/", todo.ListTodos)
		r.Post("/", todo.CreateTodo)
		r.Route("/:todoId", func(r chi.Router) {
			r.Get("/", todo.GetTodo)
			r.Put("/", todo.PutTodo)
			r.Patch("/", todo.PatchTodo)
			r.Delete("/", todo.DeleteTodo)
		})


	})

	return r
}
