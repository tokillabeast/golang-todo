package database

import (
	"github.com/tokillamockingbird/golang-todo/backend/models"
)

type BaseResponse struct {
	Deleted   int `rethinkdb:"deleted"`
	Errors    int `rethinkdb:"errors"`
	Inserted  int `rethinkdb:"inserted"`
	Replaced  int `rethinkdb:"replaced"`
	Skipped   int `rethinkdb:"skipped"`
	Unchanged int `rethinkdb:"unchanged"`
}

type InsertResponse struct {
	BaseResponse
	GeneratedKeys []string `rethinkdb:"generated_keys"`
}

type ChangesResponse struct {
	NewVal models.Todo `rethinkdb:"new_val"`
	OldVal models.Todo `rethinkdb:"old_val"`
}

type UpdateResponse struct {
	BaseResponse
	Changes []ChangesResponse `rethinkdb:"changes"`
}
