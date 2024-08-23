package handlers

import (
	"tsProject/services"

	_ "tsProject/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

type Handler struct {
	services *services.Services
}

func NewHandler(
	services services.Services,
) *Handler {
	return &Handler{
		services: &services,
	}
}

func (h *Handler) Init() {
	var jwtSecret = []byte("my_secret_key")
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allows all origins
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
	}))
	app.Post("/login", h.Login)
	app.Post("/register", h.Register)
	app.Get("/swagger/*", swagger.HandlerDefault)
	api := app.Group("/api", JWTMiddleware(jwtSecret))

	products := api.Group("/products")
	products.Get("/", h.GetAllProductsWithInstitutions)
	products.Post("/", h.AddProduct)
	products.Get("/:id", h.GetByID)
	products.Delete("/:id", h.DeleteProduct)

	api.Get("/products/:product_id/cycles/:year", h.GetCyclesByProductID)
	api.Post("/cycle", h.AddCycle)

	// Institution rotaları
	api.Post("/institution", h.AddInstitution)
	api.Get("/institution/:id", h.GetInstitutionByID)

	// Admin rotaları
	admin := api.Group("/admin")
	admin.Get("/", h.Admin)

	app.Listen(":3000")
}
