package utils

import (
	"time"
	"tsProject/config"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(userID int, username, role string) (string, error) {
	jwtKey := config.JWTKey
	expirationTime := time.Now().Add(24 * time.Hour) // Token 24 saat geçerli
	claims := &Claims{
		Username: username,
		Role:     role,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
