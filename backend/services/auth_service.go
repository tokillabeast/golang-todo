package services

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/tokillamockingbird/golang-todo/backend/models"
)

var mySigningKey = []byte("SUPER TOP SECRET") // FIXME: move to config

func GenerateToken(user models.User) (string, error) {
	/* Create the token */
	token := jwt.New(jwt.SigningMethodHS256)

	/* Set token claims */
	standardClaims := jwt.StandardClaims{}
	standardClaims.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = models.Claims{user, standardClaims}

	/* Sign the token with our secret */
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
