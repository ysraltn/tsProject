package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// JWTMiddleware creates a new middleware that protects routes using JWT
func JWTMiddleware(secret []byte) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:  secret,
		TokenLookup: "header:Authorization", // Authorization başlığından token'ı alır
		AuthScheme:  "Bearer",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
		},
	})
}

func AdminMiddleware(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	// Bir sonraki middleware veya handler'a geç
	return c.Next()
}

func SupervisorMiddleware(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	role, ok := claims["role"].(string)
	if !ok || role != "supervisor" && role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	// Bir sonraki middleware veya handler'a geç
	return c.Next()
}

// GetUserIDFromJWT retrieves the user ID from JWT token in the context
func ParseJWT(c *fiber.Ctx) (*JWTClaimsUser, error) {
	var JWTUser JWTClaimsUser
	// Token'ı kullanıcı bilgileriyle birlikte al
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// Token'ın claim'lerini map olarak al
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	// user_id'yi al ve int'e dönüştür
	userIDFloat, ok := claims["user_id"].(float64) // JSON sayıları float64 olarak gelir
	if !ok {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid user ID format")
	}

	userID := int(userIDFloat) // float64'ü int'e dönüştür
	JWTUser.UserID = userID
	JWTUser.Username, ok = claims["username"].(string)
	if !ok {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid format")
	}
	JWTUser.Role, ok = claims["role"].(string)
	if !ok {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid format")
	}
	fmt.Println("User ID:", JWTUser.UserID)
	fmt.Println("Username:", JWTUser.Username)
	fmt.Println("User Role:", JWTUser.Role)

	return &JWTUser, nil
}
