package tests

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pk-anderson/go-auth/config"
	"github.com/pk-anderson/go-auth/interfaces"
	"github.com/pk-anderson/go-auth/services"
)

func LoadTestConfig() string {
	cfg := config.LoadConfig("../config.yaml")
	return cfg.JWT.Secret
}

func TestGenerateToken(t *testing.T) {
	jwtKey := LoadTestConfig()
	authService := services.NewAuthService(jwtKey)

	id := "testId"
	email := "testuser@gmail.com"
	token, err := authService.GenerateToken(id, email)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims := &interfaces.Claims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		t.Fatalf("Expected no error when parsing token, got %v", err)
	}

	if !parsedToken.Valid {
		t.Fatal("Token should be valid")
	}

	if claims.Id != id || claims.Email != email {
		t.Fatalf("Expected claims to be {Id: %s, Email: %s}, got {Id: %s, Email: %s}", id, email, claims.Id, claims.Email)
	}
}

func TestGenerateToken_EmptyID(t *testing.T) {
	jwtKey := LoadTestConfig()
	authService := services.NewAuthService(jwtKey)

	_, err := authService.GenerateToken("", "testuser@gmail.com")
	if err == nil {
		t.Fatal("Expected an error for empty ID, got none")
	}
}

func TestGenerateToken_EmptyEmail(t *testing.T) {
	jwtKey := LoadTestConfig()
	authService := services.NewAuthService(jwtKey)

	_, err := authService.GenerateToken("123", "")
	if err == nil {
		t.Fatal("Expected an error for empty email, got none")
	}
}

func TestGenerateToken_ExpiredClaims(t *testing.T) {
	jwtKey := LoadTestConfig()
	authService := services.NewAuthService(jwtKey)

	expiredClaims := &interfaces.Claims{
		Id:    "testId",
		Email: "testuser@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, expiredClaims)
	tokenString, _ := token.SignedString([]byte(jwtKey))

	claims, err := authService.ValidateToken(tokenString)

	if err == nil {
		t.Fatal("Expected error for expired token, got none")
	}

	if claims != nil {
		t.Fatal("Expected claims to be nil for expired token")
	}
}
