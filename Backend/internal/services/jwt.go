package services

import (
	"errors"
	"strings"
	"time"

	"OnlineBar/Backend/internal/models"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(username string, id string) (string, error) {

	var sampleSecretKey = []byte("OnlineBar")

	claims := models.Claims{
		Authorized: true,
		Exp:        time.Now().Add(1 * time.Hour).Unix(),
		User:       username,
		ID:         id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(sampleSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func CheckJWT(token string, claims *models.Claims) error {
	const secretKey = "OnlineBar"

	if !strings.HasPrefix(token, "Bearer ") {
		return errors.New("invalid token format, missing Bearer prefix")
	}

	token = strings.TrimPrefix(token, "Bearer ")

	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return err
	}

	if err := claims.Valid(); err != nil {
		return errors.New("token not valid")
	}

	if claims.ExpiresAt > time.Now().Unix() {
		return errors.New("token has expired")
	}

	return nil
}
