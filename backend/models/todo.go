package models

type Todo struct {
	Id     string `gorethink:"id,omitempty" json:"id,omitempty"`
	Text   string `gorethink:"text,omitempty" json:"text,omitempty"`
	Status string `gorethink:"status,omitempty" json:"status,omitempty"`
	UserId string `gorethink:"user_id,omitempty" json:"id,omitempty"`
	Base
}

type Todos []Todo
