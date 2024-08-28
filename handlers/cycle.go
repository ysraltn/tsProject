package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Summary Add Cycle
// @Description Add Product Cycle
// @Tags cycles
// @Accept  json
// @Produce  json
// @Param product body models.Cycle true "Cycle data"
// @Success 200 {object} models.Cycle
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/cycle [post]
func (h *Handler) AddCycle(c *fiber.Ctx) error {
	var cycle Cycle
	if err := c.BodyParser(&cycle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}
	err := h.services.CycleService.Add(
		cycle.ProductID,
		cycle.Year,
		cycle.Month,
		cycle.CycleNo,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add cycle",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Cycle added successfully", cycle)
}

// @Summary Get Cycles by Product ID and Year
// @Description Get all cycles for a specific product in a given year
// @Tags cycles
// @Accept  json
// @Produce  json
// @Param product_id path int true "Product ID"
// @Param year path int true "Year"
// @Success 200 {array} models.Cycle
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /api/products/{product_id}/cycles/{year} [get]
func (h *Handler) GetCyclesByProductID(c *fiber.Ctx) error {
	productIDParam := c.Params("product_id")
	productID, err := strconv.Atoi(productIDParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}
	yearParam := c.Params("year")
	year, err := strconv.Atoi(yearParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}
	cycles, err := h.services.CycleService.GetProductsCyclesByYear(productID, year)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch cycles",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Cycles fetched successfully", cycles)
}
