package database

import (
	r "gopkg.in/gorethink/gorethink.v3"

	e "github.com/tokillamockingbird/golang-todo/backend/errors"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

const UserTable = "users" // FIXME: is it a best way to store table name?

func GetUsers() models.Users { // FIXME: add filtering & order support
	users := models.Users{}
	response, err := r.Table(UserTable).Run(Connect())
	e.CheckAndLogError(err)
	err = response.All(&users)
	e.CheckAndLogError(err)
	return users
}

func GetUser(id string) models.User {
	var user models.User
	response, err := r.Table(UserTable).Run(Connect())
	e.CheckAndLogError(err)
	err = response.One(&user)
	e.CheckAndLogError(err)
	return user
}

func CreateUser(user models.User) models.User {
	return models.User{}
}
