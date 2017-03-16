package database

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"

	"github.com/tokillamockingbird/golang-todo/backend/config"
)

func Connect() *r.Session {
	conf := config.ReadConfig(config.DevelopmentEnv)
	session, err := r.Connect(r.ConnectOpts{
		Address:    conf.DatabaseUrl,
		Database:   conf.DatabaseName,
		InitialCap: conf.DatabaseInitialCap,
		MaxOpen:    conf.DatabaseMaxOpen,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}
	return session
}
