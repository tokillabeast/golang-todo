package database

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"

	"github.com/tokillamockingbird/golang-todo/backend/models"
)

const todoTable = "test" // FIXME: is it a best way to store table name?

func RepoListTodo() models.Todos {
	todos := models.Todos{}
	response, err := r.Table(todoTable).Run(Connect())
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
	response, err := r.Table(todoTable).Get(id).Run(Connect())
	if err != nil {
		log.Fatalln(err)
	}
	err = response.One(&todo) // Check if result return values
	if err != nil {
		log.Fatalln(err)
	}
	return todo
}

func RepoCreateTodo(todo models.Todo) models.Todo {
	response, err := r.Table(todoTable).Insert(todo).RunWrite(Connect())
	if err != nil {
		log.Fatalln(err)
	}
	if len(response.GeneratedKeys) != 1 {
		log.Fatalln("GeneratedKeys doesn't contain 1 element")
	}
	todo.Id = response.GeneratedKeys[0]
	return todo
}

func RepoUpdateTodo(todo models.Todo) models.Todo {
	err := r.Table(todoTable).Replace(todo).Exec(Connect())
	if err != nil {
		log.Fatalln(err)
	}
	return todo
}

func RepoPatchTodo(todoId string, todo models.Todo) models.Todo {
	updateOpt := r.UpdateOpts{ReturnChanges: true}
	response, err := r.Table(todoTable).Get(todoId).Update(todo, updateOpt).RunWrite(Connect())
	if err != nil {
		log.Fatalln(err)
	}
	if len(response.Changes) != 1 {
		log.Fatalln("Changes doesn't contain 1 element")
	}
	newValue := response.Changes[0].NewValue.(map[string]interface{})

	return models.Todo{
		Id:     newValue["id"].(string),
		Text:   newValue["text"].(string),
		Status: newValue["status"].(string),
	}

}

func RepoDeleteTodo(id string) {
	err := r.Table(todoTable).Get(id).Delete().Exec(Connect())
	if err != nil {
		log.Fatalln(err)
	}
}
