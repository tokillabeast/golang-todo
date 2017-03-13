package main

import (
	"log"
	"net/http"

	"github.com/pressly/chi"
	r "gopkg.in/gorethink/gorethink.v3"
)

func RethinkTest() string {
	session, err := r.Connect(r.ConnectOpts{
		Address: "golang-todo_rethinkdb",
	})
	if err != nil {
		log.Fatalln(err)
	}
	res, err := r.Expr("Hello World!").Run(session)
	if err != nil {
		log.Fatalln(err)
	}

	var response string
	err = res.One(&response)
	if err != nil {
		log.Fatalln(err)
	}
	return response
}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result := RethinkTest()
		w.Write([]byte(result))
	})
	http.ListenAndServe("0.0.0.0:8000", r)
}
