package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pk-anderson/go-auth/interfaces"
)

type AuthService struct {
	jwtKey []byte
}

func (s *AuthService) GenerateToken(id, email string) (string, error) {
	if id == "" {
		return "", errors.New("ID cannot be empty")
	}
	if email == "" {
		return "", errors.New("email cannot be empty")
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &interfaces.Claims{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *AuthService) ValidateToken(tokenStr string) (*interfaces.Claims, error) {
	claims := &interfaces.Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func NewAuthService(jwtKey string) *AuthService {
	return &AuthService{
		jwtKey: []byte(jwtKey),
	}
}
