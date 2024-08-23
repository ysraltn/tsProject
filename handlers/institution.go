package handlers

import (
	"strconv"
	"tsProject/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Add Institution
// @Description Register an institution
// @Tags institutions
// @Accept  json
// @Produce  json
// @Param product body models.Institution true "Institution data"
// @Success 200 {object} models.Institution
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/institution [post]
func (h *Handler) AddInstitution(c *fiber.Ctx) error {
	var institution models.Institution
	if err := c.BodyParser(&institution); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}
	err := h.services.InstitutionService.Add(
		institution.Name,
		institution.City,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add product",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Institution added successfully", institution)
}

// @Summary Get Institution by ID
// @Description Get a single institution by its ID
// @Tags institutions
// @Accept  json
// @Produce  json
// @Param id path int true "Institution ID"
// @Success 200 {object} models.Institution
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/institution/{id} [get]
func (h *Handler) GetInstitutionByID(c *fiber.Ctx) error {
	// ID parametresini al ve integer'a çevir
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid institution ID",
			"error":   err.Error(),
		})
	}

	// Veritabanından ID'ye göre kurumu al
	institution, err := h.services.InstitutionService.GetByID(id)
	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch institution",
			"error":   err.Error(),
		})
	}

	// Kurum bulunduysa, JSON olarak döndür
	return Response(c, 200, "Institution fetched successfully", institution)
}
