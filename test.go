package main

import (
	"net/http"

	"github.com/pressly/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe("0.0.0.0:8000", r)
}
