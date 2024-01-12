package models

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Authorized bool   `json:"authorized"`
	Exp        int64  `json:"exp"`
	User       string `json:"user"`
	ID         string `json:"id"`
	jwt.StandardClaims
}
