package tests

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pk-anderson/go-auth/interfaces"
	"github.com/pk-anderson/go-auth/services"
)

func TestValidateToken(t *testing.T) {
	jwtKey := LoadTestConfig()
	authService := services.NewAuthService(jwtKey)

	id := "testId"
	email := "testuser@gmail.com"
	token, _ := authService.GenerateToken(id, email)

	claims, err := authService.ValidateToken(token)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.Id != id || claims.Email != email {
		t.Fatalf("Expected claims to be {Id: %s, Email: %s}, got {Id: %s, Email: %s}", id, email, claims.Id, claims.Email)
	}
}

func TestValidateToken_InvalidToken(t *testing.T) {
	jwtKey := LoadTestConfig()
	authService := services.NewAuthService(jwtKey)

	claims, err := authService.ValidateToken("invalid.token.string")

	if err == nil {
		t.Fatal("Expected error for invalid token, got none")
	}

	if claims != nil {
		t.Fatal("Expected claims to be nil for invalid token")
	}
}

func TestValidateToken_ExpiredToken(t *testing.T) {
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
