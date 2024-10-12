package interfaces

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
