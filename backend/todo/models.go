package todo

import "time"

type Base struct {
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

type Todo struct {
	Id     string `json:"id",omitempty`
	Text   string `json:"text"`
	Status string `json:"status"`
	Base
}

type Todos []Todo


