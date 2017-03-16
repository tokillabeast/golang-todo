package models

import "time"

type Base struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type Todo struct {
	Id     string `gorethink:"id,omitempty" json:"id,omitempty"`
	Text   string `gorethink:"text" json:"text"`
	Status string `gorethink:"text" json:"status"`
	Base
}

type Todos []Todo
