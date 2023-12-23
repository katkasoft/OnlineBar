package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var sampleSecretKey = []byte("OnlineBar")

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour)
	claims["authorized"] = true
	claims["user"] = username

	tokenString, err := token.SignedString(sampleSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
