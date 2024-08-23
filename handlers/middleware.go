package handlers

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// JWTMiddleware creates a new middleware that protects routes using JWT
func JWTMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: secret,
	})
}
