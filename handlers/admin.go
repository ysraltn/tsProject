package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type AddUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// Admin handler
// @Summary Admin access
// @Description This endpoint is restricted to admin users only. It validates the JWT token and checks if the user has an "admin" role.
// @Tags admin
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} string "Welcome Admin!"
// @Security BearerAuth
// @Router /api/admin [get]
func (h *Handler) Admin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	role := claims["role"].(string)

	if role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Access denied"})
	}

	return c.SendString("Welcome Admin!")
}

// @Summary AddUser
// @Description Add user with given role
// @Tags admin
// @Accept  json
// @Produce  json
// @Param loginRequest body AddUserRequest true "Add User Request"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]string "Login Error"
// @Security BearerAuth
// @Router /api/admin/user [post]
func (h *Handler) AddUser(c *fiber.Ctx) error {
	var addUserRequest AddUserRequest
	err := c.BodyParser(&addUserRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}
	err = h.services.UserService.Add(addUserRequest.Username, addUserRequest.Password, addUserRequest.Role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error adding user",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "user added successfully", addUserRequest)
}

// @Summary GetAllUsers
// @Description Get all users
// @Tags admin
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /api/admin/user [get]
func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.services.UserService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error getting users",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "users fetched successfully", users)
}
