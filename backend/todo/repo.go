package todo

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"

	"github.com/tokillamockingbird/golang-todo/backend/database"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

const todoTable = "test" // FIXME: is it a best way to store table name?

func RepoListTodo() models.Todos {
	todos := models.Todos{}
	session := database.Connect()
	response, err := r.Table(todoTable).Run(session)
	if err != nil {
		log.Fatalln(err)
	}
	err = response.All(&todos)
	if err != nil {
		log.Fatalln(err)
	}
	return todos
}

func RepoFindTodo(id string) models.Todo {
	todo := models.Todo{}
	session := database.Connect()
	response, err := r.Table(todoTable).Get(id).Run(session)
	if err != nil {
		log.Fatalln(err)
	}
	err = response.One(&todo)
	if err != nil {
		log.Fatalln(err)
	}
	return todo
}

func RepoCreateTodo(todo models.Todo) models.Todo { // FIXME: test, looks like doesn't work correct
	var result database.InsertResponse
	session := database.Connect()
	response, err := r.Table(todoTable).Insert(todo).Run(session)
	if err != nil {
		log.Fatalln(err)
	}
	err = response.One(&result)
	if err != nil {
		log.Fatalln(err)
	}
	todo.Id = result.GeneratedKeys[0]
	return todo
}

func RepoUpdateTodo(todo models.Todo) models.Todo {
	session := database.Connect()
	err := r.Table(todoTable).Replace(todo).Exec(session)
	if err != nil {
		log.Fatalln(err)
	}
	return todo
}

func RepoPatchTodo(todoId string, json interface{}) models.Todo {
	var result database.UpdateResponse
	session := database.Connect()
	response, err := r.Table(todoTable).Get(todoId).Update(json).Run(session)
	if err != nil {
		log.Fatalln(err)
	}
	err = response.One(&result)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Changes[0].NewVal
}

func RepoDeleteTodo(id string) {
	session := database.Connect()
	err := r.Table(todoTable).Get(id).Delete().Exec(session)
	if err != nil {
		log.Fatalln(err)
	}
}
