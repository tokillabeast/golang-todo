package database

import (
	"errors"
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

func GetAuthenticateUser(authenticate models.AuthenticateRequest) (models.User, error) { // FIXME: use GetUsers?
	var users []models.User
	response, err := r.Table(UserTable).Filter(authenticate).Run(Connect())
	if err != nil {
		return models.User{}, errors.New("Internal error")
	}
	if response == nil {
		return models.User{}, errors.New("Internal error")
	}
	err = response.All(&users)
	if err != nil {
		return models.User{}, errors.New("Internal error")
	}
	switch len(users) {
	case 0:
		return models.User{}, errors.New("User not found")
	case 1:
		return users[0], nil
	default:
		return models.User{}, errors.New("Internal error: Several users found")
	}
}

func GetUser(id string) models.User {
	var user models.User
	response, err := r.Table(UserTable).Run(Connect())
	e.CheckAndLogError(err)
	err = response.One(&user)
	e.CheckAndLogError(err)
	return user
}

func CreateUser(user models.User) (models.User, error) {
	user.Created, user.Modified = time.Now(), time.Now()
	response, err := r.Table(UserTable).Insert(user).RunWrite(Connect())
	if err != nil {
		return models.User{}, err
	}
	if response.Inserted != 1 && len(response.GeneratedKeys) != 1 {
		return models.User{}, errors.New("GeneratedKeys doesn't contain 1 element")
	}
	user.Id = response.GeneratedKeys[0]
	return user, nil
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
