package models

import (
	"time"
)

type Base struct {
	Created  time.Time `gorethink:"created,omitempty" json:"created,omitempty"`
	Modified time.Time `gorethink:"modified,omitempty" json:"modified,omitempty"`
}

type Todo struct {
	Id     string `gorethink:"id,omitempty" json:"id,omitempty"`
	Text   string `gorethink:"text,omitempty" json:"text,omitempty"`
	Status string `gorethink:"status,omitempty" json:"status,omitempty"`
	Base
}

type Todos []Todo
