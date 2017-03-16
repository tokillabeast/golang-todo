package main

import (
	"net/http"

	"github.com/tokillamockingbird/golang-todo/backend/config"
)

func main() {
	r := Routes()
	config := config.ReadConfig(config.DevelopmentEnv) // FIXME: config.DevelopmentEnv looks wrong
	http.ListenAndServe(config.BaseUrl, r)
}
