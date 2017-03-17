package models

type User struct {
	Username string `gorethink:"username" json:"username"` // FIXME: validation, unique check
	Email    string `gorethink:"email" json"email"`        // FIXME: validation, unique check
	Password string `gorethink:"password" json"password"`  // FIXME: encryption
	Base
}

type Users []User
