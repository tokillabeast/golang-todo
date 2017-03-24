package models

import "time"

type Base struct {
	Created  time.Time `gorethink:"created,omitempty" json:"created,omitempty"`
	Modified time.Time `gorethink:"modified,omitempty" json:"modified,omitempty"`
}

type AuthenticateRequest struct {
	Email    string `gorethink:"email,omitempty" json:"email,omitempty"`
	Username string `gorethink:"username,omitempty" json:"username,omitempty"`
	Password string `gorethink:"password" json:"password"`
}

type AuthenticateResponse struct {
	User
	Token string `json:"token"`
}
