package handlers

import (
	"tsProject/config"
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
	jwtSecret := config.JWTKey
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allows all origins
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Static("/", "./public")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/profile.html") // HTML dosyasını sunar
	})
	app.Post("/login", h.Login)
	app.Post("/register", h.Register)
	app.Get("/swagger/*", swagger.HandlerDefault)
	
	api := app.Group("/api", JWTMiddleware(jwtSecret))

	products := api.Group("/products")
	products.Use(SupervisorMiddleware)
	products.Get("/yok", h.GetAllProductsWithInstitutionAndCycle)
	products.Get("/download", h.DownloadProductsWithInstitutionAndCycle)
	products.Get("/", h.GetAllProductsWithInstitutions)
	products.Post("/", h.AddProduct)
	products.Get("/:id", h.GetByID)
	products.Delete("/:id", h.DeleteProduct)
	products.Get("/filterbyinsid/:id", h.GetProductsByInstitutionID)
	products.Post("/filter", h.FilterProducts)

	api.Get("/products/:product_id/cycles/:year", h.GetCyclesByProductID)
	api.Post("/cycle", h.AddCycle)

	// Institution rotaları
	api.Get("/institution", h.GetAllInstitutions)
	api.Post("/institution", h.AddInstitution)
	api.Get("/institution/:id", h.GetInstitutionByID)

	profile := api.Group("/profile")
	profile.Get("/", h.GetProfile)
	profile.Get("/institutions", h.GetAssignedProductsHandler)
	profile.Get("/user", h.GetAllForUsers)

	// Admin rotaları
	admin := api.Group("/admin")
	admin.Use(AdminMiddleware)
	admin.Get("/", h.Admin)
	admin.Post("/user", h.AddUser)
	admin.Put("/user", h.UpdateUserByID)
	admin.Get("/user", h.GetAllUsers)

	app.Listen(":3000")
}
