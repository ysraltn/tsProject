package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary Login
// @Description Authenticate user and return a JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param loginRequest body LoginRequest true "Login Request"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]string "Login Error"
// @Router /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	var loginRequest LoginRequest
	err := c.BodyParser(&loginRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}

	token, userRole, err := h.services.AuthService.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Login Error",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"token":   token,
		"role":    userRole,
	})
}

// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param registerRequest body RegisterRequest true "Register Request"
// @Success 200 {object} map[string]string "user registered successfully"
// @Failure 400 {object} map[string]string "error registering user"
// @Router /register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
	var registerRequest RegisterRequest
	err := c.BodyParser(&registerRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}
	err = h.services.AuthService.Register(registerRequest.Username, registerRequest.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error registering user",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "user registered successfully", registerRequest)
}
