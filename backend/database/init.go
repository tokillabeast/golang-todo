package database

import (
	r "gopkg.in/gorethink/gorethink.v3"

	"github.com/tokillamockingbird/golang-todo/backend/config"
	"github.com/tokillamockingbird/golang-todo/backend/errors"
	"log"
)

func InitDatabase(configuration config.Config) {
	var dbExist bool
	response, err := r.DBList().Contains(configuration.DatabaseName).Run(Connect())
	errors.CheckAndLogError(err)
	response.One(&dbExist)
	if !dbExist {
		err = r.DBCreate(configuration.DatabaseName).Exec(Connect())
		errors.CheckAndLogError(err)
	}

	for _, tableName := range []string{TodoTable, UserTable} {
		var tableExist bool
		response, err = r.TableList().Contains(tableName).Run(Connect())
		errors.CheckAndLogError(err)
		response.One(&tableExist)
		if !tableExist {
			err := r.TableCreate(tableName).Exec(Connect())
			errors.CheckAndLogError(err)
		}
	}
}
