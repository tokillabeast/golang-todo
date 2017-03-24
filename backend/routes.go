package main

import (
	"net/http"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"

	"github.com/tokillamockingbird/golang-todo/backend/handlers"
	"github.com/tokillamockingbird/golang-todo/backend/routes"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Golang ;)"))
	})
	r.Get("/api/v1/status", handlers.NotImplemented)

	r.Post("/api/v1/register", handlers.RegisterNewUser)
	r.Post("/api/v1/authenticate", handlers.AuthenticateUser)

	r.Mount("/api/v1/todos", routes.TodoResource{}.Routes())
	r.Mount("/api/v1/users", routes.UsersResource{}.Routes())

	return r
}
