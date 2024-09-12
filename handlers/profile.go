package handlers

import (
	"tsProject/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get User profile
// @Description Get the profile information of the authenticated user
// @Tags profile
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Security BearerAuth
// @Router /api/profile [get]
func (h *Handler) GetProfile(c *fiber.Ctx) error {
	user, err := ParseJWT(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid token: missing userID",
		})
	}

	// // Kullanıcı adını güvenli bir şekilde çek
	// username, ok := claims["username"].(string)
	// if !ok {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Invalid token: missing username",
	// 	})
	// }

	// // Kullanıcı rolünü güvenli bir şekilde çek
	// role, ok := claims["role"].(string)
	// if !ok {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Invalid token: missing role",
	// 	})
	// }

	// Kullanıcı bilgilerini JSON formatında döndürüyoruz
	var userResponse models.UserResponse
	userResponse.ID = user.UserID
	userResponse.Role = user.Role
	userResponse.Username = user.Username
	return Response(c, 200, "User profile fetched successfully", userResponse)
}

// @Summary Get Users Names
// @Description Get users names for select input
// @Tags profile
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserForUsers
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Security BearerAuth
// @Router /api/profile/user [get]
func (h *Handler) GetAllForUsers(c *fiber.Ctx) error {
	users, err := h.services.UserService.GetAllForUsers()
	if err != nil {
		return err
	}
	return Response(c, 200, "Users fetched successfully", users)
}
