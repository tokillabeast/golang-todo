package middleware

import (
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
)

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("My Super Secret Key"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
	// Setup other options: errorHandler, etc... - https://github.com/auth0/go-jwt-middleware#options
})
