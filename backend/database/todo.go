package database

import (
	"log"
	"time"

	r "gopkg.in/gorethink/gorethink.v3"

	e "github.com/tokillamockingbird/golang-todo/backend/errors"
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

const TodoTable = "todos" // FIXME: is it a best way to store table name?

func GetTodos() models.Todos {
	todos := models.Todos{}
	response, err := r.Table(TodoTable).Run(Connect())
	e.CheckAndLogError(err)
	err = response.All(&todos)
	e.CheckAndLogError(err)
	return todos
}

func GetTodo(id string) models.Todo {
	todo := models.Todo{}
	response, err := r.Table(TodoTable).Get(id).Run(Connect())
	e.CheckAndLogError(err)
	err = response.One(&todo) // Check if result return values
	e.CheckAndLogError(err)
	return todo
}

func CreateTodo(todo models.Todo) models.Todo {
	todo.Created = time.Now()
	todo.Modified = time.Now()
	response, err := r.Table(TodoTable).Insert(todo).RunWrite(Connect())
	e.CheckAndLogError(err)
	if len(response.GeneratedKeys) != 1 {
		log.Fatalln("GeneratedKeys doesn't contain 1 element")
	}
	todo.Id = response.GeneratedKeys[0]
	return todo
}

func UpdateTodo(todo models.Todo) models.Todo { // FIXME: leave only update or patch
	todo.Modified = time.Now()
	err := r.Table(TodoTable).Replace(todo).Exec(Connect())
	e.CheckAndLogError(err)
	return todo
}

func PatchTodo(todoId string, todo models.Todo) models.Todo {
	todo.Modified = time.Now()
	updateOpt := r.UpdateOpts{ReturnChanges: true}
	response, err := r.Table(TodoTable).Get(todoId).Update(todo, updateOpt).RunWrite(Connect())
	e.CheckAndLogError(err)
	if len(response.Changes) != 1 { // FIXME: don't need to check Changes, we can send PATCH request with current value
		log.Fatalln("Changes doesn't contain 1 element")
	}
	newValue := response.Changes[0].NewValue.(map[string]interface{})
	// FIXME: get latest version of user

	return models.Todo{
		Id:     newValue["id"].(string),
		Text:   newValue["text"].(string),
		Status: newValue["status"].(string),
	}

}

func DeleteTodo(todoId string) {
	err := r.Table(TodoTable).Get(todoId).Delete().Exec(Connect())
	e.CheckAndLogError(err)
}
