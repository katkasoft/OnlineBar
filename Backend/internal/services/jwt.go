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

	// Проверка префикса "Bearer"
	if !strings.HasPrefix(token, "Bearer ") {
		return errors.New("invalid token format, missing Bearer prefix")
	}

	// Извлечение токена без префикса "Bearer"
	token = strings.TrimPrefix(token, "Bearer ")

	// Парсинг токена с заполнением структуры Claims
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		// Возвращение ошибки, если парсинг токена не удался
		return err
	}

	// Проверка валидности токена
	if err := claims.Valid(); err != nil {
		// Возвращение ошибки, если токен не валиден
		return errors.New("token not valid")
	}

	// Токен валиден
	return nil
}
