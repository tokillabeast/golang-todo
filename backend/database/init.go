package database

import (
	r "gopkg.in/gorethink/gorethink.v3"

	"github.com/tokillamockingbird/golang-todo/backend/config"
	"github.com/tokillamockingbird/golang-todo/backend/errors"
	"fmt"
)

func DatabaseInitialization(configuration config.Config) {
	var databases []string
	response, err := r.DBList().Run(Connect())
	errors.CheckAndLogError(err)
	response.All(&databases)
	dbExist := false
	for _, databaseName := range databases {
		if databaseName == configuration.DatabaseName {
			dbExist = true
			break
		}
	}
	fmt.Println(dbExist)
	if !dbExist {
		err = r.DBCreate(configuration.DatabaseName).Exec(Connect())
		errors.CheckAndLogError(err)
	}

	for _, tableName := range []string{TodoTable, UserTable} {
		var tables []string
		tableExist := false
		response, err = r.TableList().Run(Connect())
		errors.CheckAndLogError(err)
		for _, table := range tables {
			if tableName == table {
				tableExist = true
				break
			}
		}
		if !tableExist {
			err := r.TableCreate(tableName).Exec(Connect())
			errors.CheckAndLogError(err)
		}
	}
}
