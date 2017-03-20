package main

import (
	"net/http"

	"github.com/tokillamockingbird/golang-todo/backend/config"
	"github.com/tokillamockingbird/golang-todo/backend/database"
)

func main() {
	r := Routes()
	configuration := config.ReadConfig(config.DevelopmentEnv) // FIXME: config.DevelopmentEnv looks wrong
	database.InitDatabase(configuration)
	http.ListenAndServe(configuration.BaseUrl, r)
}
