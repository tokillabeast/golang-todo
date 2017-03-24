package database

import (
	"testing"

	r "gopkg.in/gorethink/gorethink.v3"
	"github.com/stretchr/testify/assert"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	session := Connect()
	r.DBCreate(TestDatabaseName).Exec(session)
	defer r.DBDrop(TestDatabaseName).Exec(session)
	r.TableCreate(UserTable).Exec(session)
	defer r.TableDrop(UserTable).Exec(session)
	user := models.User{Username: "username", Email:"email@test.com", Password:"password"}
	user, err := CreateUser(user)
	assert.NoError(err)
	assert.NotNil(user.Id)
}
