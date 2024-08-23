package handlers

import (
	"strconv"
	"tsProject/models"

	"github.com/gofiber/fiber/v2"
)

// @Summary Add Product
// @Description Register a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body models.Product true "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/products [post]
func (h *Handler) AddProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}
	err := h.services.ProductService.Add(
		product.Serial,
		product.Brand,
		product.Model,
		product.Responsible,
		product.Owner,
		product.Status,
		product.InstitutionID,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to add product",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Products added successfully", product)
}

func (h *Handler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.services.ProductService.GetAll()
	if err != nil {
		return err
	}
	return Response(c, 200, "Products fetched successfully", products)
}

// @Summary Get All Products with institution names
// @Description Get all products with institution names
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProductWithInstitution
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /api/products [get]
func (h *Handler) GetAllProductsWithInstitutions(c *fiber.Ctx) error {
	products, err := h.services.ProductService.GetAllWithInstitutions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch products with institutions",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Products with institutions fetched successfully", products)
}

// @Summary Delete Product
// @Description Delete a product by ID
// @Tags products
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/products/{id} [delete]
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}
	err = h.services.ProductService.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get product",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Product deleted successfully", nil)
}

// @Summary Get Product by ID
// @Description Get a product by its ID
// @Tags products
// @Param id path int true "Product ID"
// @Produce json
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /api/products/{id} [get]
func (h *Handler) GetByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}
	product, err := h.services.ProductService.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch product",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Product fetched successfully", product)
}
