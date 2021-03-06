package routes

import (
	"github.com/pressly/chi"

	"github.com/tokillamockingbird/golang-todo/backend/handlers"
)

type UsersResource struct{}

func (rs UsersResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
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
