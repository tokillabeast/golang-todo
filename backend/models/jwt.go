package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	User
	jwt.StandardClaims
}
