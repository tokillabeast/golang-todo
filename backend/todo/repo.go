package todo

import (
	"log"

	r "gopkg.in/gorethink/gorethink.v3"

	"github.com/tokillamockingbird/golang-todo/backend/database"
)

func RepoListTodo() Todos{
	todos := Todos{}
	session := database.Connect()
	response, err := r.Table("test").Run(session) // FIXME: "test"
	if err != nil {
		log.Fatalln(err)
	}
	err = response.All(&todos)
	if err != nil {
		log.Fatalln(err)
	}
	return todos
}

func RepoFindTodo(id string) Todo {
	todo := Todo{}
	session := database.Connect()
	response, err := r.Table("test").Get(id).Run(session) // FIXME: "test"
	if err != nil {
		log.Fatalln(err)
	}
	err = response.One(&todo)
	if err != nil {
		log.Fatalln(err)
	}
	return todo
}

func RepoCreateTodo(todo Todo) Todo {
	var result database.InsertRespose
	session := database.Connect()
	response, err := r.Table("test").Insert(todo).Run(session)
	if err != nil {
		log.Fatalln(err)
	}
	err = response.One(&result)
	if err != nil {
		log.Fatalln(err)
	}
	todo.Id = result.Generated_keys[0]
	return todo
}

func RepoDeleteTodo(id string) {
	session := database.Connect()
	err := r.Table("test").Get(id).Delete().Exec(session)
	if err != nil {
		log.Fatalln(err)
	}
}
