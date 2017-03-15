package todo

import (
	"log"
	"fmt"

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

func RepoCreateTodo(t Todo) {
	session := database.Connect()
	todoItem := Todo{Text: "TodoItem Test", Status: "Active"}
	err := r.Table("test").Insert(todoItem).Exec(session)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Completed!!!!!")
}

func RepoDeleteTodo(id int) {

}
