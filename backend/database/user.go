package database

import (
	"log"
	"time"

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
	user.Created = time.Now()
	user.Modified = time.Now()
	response, err := r.Table(UserTable).Insert(user).RunWrite(Connect())
	e.CheckAndLogError(err)
	if len(response.GeneratedKeys) != 1 {
		log.Fatalln("GeneratedKeys doesn't contain 1 element")
	}
	user.Id = response.GeneratedKeys[0]
	return user
}

func UpdateUser(user models.User) models.User {
	user.Modified = time.Now()
	err := r.Table(UserTable).Replace(user).Exec(Connect())
	e.CheckAndLogError(err)
	return user
}

func PatchUser(userId string, user models.User) models.User {
	user.Modified = time.Now()
	err := r.Table(UserTable).Get(userId).Update(user).Exec(Connect())
	e.CheckAndLogError(err)
	response, err := r.Table(UserTable).Get(userId).Run(Connect())
	response.One(&user)
	return user
}

func DeleteUser(userId string) {
	err := r.Table(UserTable).Get(userId).Delete().Exec(Connect())
	e.CheckAndLogError(err)
}
