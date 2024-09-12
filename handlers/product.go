package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
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
		product.Status,
		product.InstitutionID,
		product.ResponsibleID,
		product.OwnerID,
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

// @Summary Get All Products with institution and cycle
// @Description Get all products with institution and cycle
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProductWithInstitutionAndCycle
// @Failure 400 {object} map[string]string
// @Security BearerAuth
// @Router /api/products/yok [get]
func (h *Handler) GetAllProductsWithInstitutionAndCycle(c *fiber.Ctx) error {
	products, err := h.services.ProductService.GetAllProductsWithInstitutionAndCycle()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch products with institutions and cycles",
			"error":   err.Error(),
		})
	}
	return Response(c, http.StatusOK, "Products with institutions and cycles fetched successfully", products)
}

// @Summary Get Assigned Products
// @Description Get products assigned to the current user based on their JWT token
// @Tags profile
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProductWithInstitutionAndCycle
// @Failure 400 {object} map[string]string "Invalid user ID"
// @Failure 500 {object} map[string]string "Failed to fetch assigned products"
// @Security BearerAuth
// @Router /api/profile/institutions [get]
func (h *Handler) GetAssignedProductsHandler(c *fiber.Ctx) error {
	// JWT'den kullanıcı ID'sini al
	user, err := ParseJWT(c)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user",
			"error":   err.Error(),
		})
	}

	// GetAssignedProducts fonksiyonunu çağır
	products, err := h.services.ProductService.GetAssignedProducts(user.UserID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch assigned products",
			"error":   err.Error(),
		})
	}

	// Başarıyla dönen ürünleri JSON olarak yanıtla
	return Response(c, http.StatusOK, "Assigned products fetched successfully", products)
}

// @Summary Filter Products by Institution ID
// @Description Get products by institution ID
// @Tags products
// @Param id path int true "Institution ID"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ProductWithInstitutionAndCycle
// @Failure 400 {object} map[string]string "Invalid institution ID"
// @Failure 500 {object} map[string]string "Failed to fetch filtered products"
// @Security BearerAuth
// @Router /api/products/filterbyinsid/{id} [get]
func (h *Handler) GetProductsByInstitutionID(c *fiber.Ctx) error {
	institutionIDParam := c.Params("id")
	institutionID, err := strconv.Atoi(institutionIDParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   err.Error(),
		})
	}
	products, err := h.services.ProductService.FilterByInstitutionID(institutionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch products",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Products fetched successfully", products)
}

//databasede kurumlar tekrarlı kaydedilmiş onu düzelt!!!!!!!!!!

// @Summary Filter Products
// @Description Filter products by given filter
// @Tags products
// @Accept  json
// @Produce  json
// @Param filter body models.ProductFilter true "Product filter"
// @Success 200 {array} models.ProductWithInstitutionAndCycle
// @Failure 400 {object} map[string]string "Invalid JSON format"
// @Failure 500 {object} map[string]string "Failed to fetch products"
// @Security BearerAuth
// @Router /api/products/filter [post]
func (h *Handler) FilterProducts(c *fiber.Ctx) error {
	var filter models.ProductFilter
	if err := c.BodyParser(&filter); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid JSON format",
			"error":   err.Error(),
		})
	}
	products, err := h.services.ProductService.Filter(filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch products",
			"error":   err.Error(),
		})
	}
	return Response(c, 200, "Products fetched successfully", products)
}

func (h *Handler) DownloadProductsWithInstitutionAndCycle(c *fiber.Ctx) error {
	format := c.Query("format")
	fmt.Println("format", format)
	excelBuffer, err := h.services.ProductService.DownloadCycles(format)
	fmt.Println("excelBuffer", excelBuffer)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to download",
			"error":   err.Error(),
		})
	}
	fileName := "donguler_" + time.Now().Format("20060102_150405") + ".xlsx" // Örnek: products_20240911_153045.xlsx

	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename="+fileName)
	c.Set("Content-Transfer-Encoding", "binary")
	return c.SendStream(excelBuffer)

}
