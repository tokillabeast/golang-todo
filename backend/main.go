package main

import (
	"net/http"

	"github.com/tokillamockingbird/golang-todo/backend/config"
	"github.com/tokillamockingbird/golang-todo/backend/route"
)

func main() {
	r := route.Routes()
	config := config.ReadConfig(config.DevelopmentEnv) // FIXME: config.DevelopmentEnv looks wrong
	http.ListenAndServe(config.BaseUrl, r)
}
