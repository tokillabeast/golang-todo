package main

import (
	"log"
	"net/http"

	"github.com/pressly/chi"
	r "gopkg.in/gorethink/gorethink.v3"

	"github.com/tokillamockingbird/golang-todo/backend/config"
)

func RethinkTest() string {
	conf := config.ReadConfig(config.DevelopmentEnv)
	session, err := r.Connect(r.ConnectOpts{
		Address: conf.DatabaseUrl,
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
	conf := config.ReadConfig(config.DevelopmentEnv)
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result := RethinkTest()
		w.Write([]byte(result))
	})
	http.ListenAndServe(conf.BaseUrl, r)
}
