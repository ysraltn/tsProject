package main

import (
	"fmt"
	"log"
	"tsProject/config"
	"tsProject/database"
	_ "tsProject/docs"
	"tsProject/handlers"
	"tsProject/logger"
	"tsProject/models"
	"tsProject/repositories"
	"tsProject/services"
	"tsProject/utils"
)

// @title Fiber API Documentation
// @version 1.0
// @description This is a sample API with JWT authentication.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger := logger.NewLogger("logs.log")
	config.Load()
	db, err := database.Init()
	if err != nil {
		logger.Error(models.ErrorLog{
			Message: "Failed to connect database",
			Error:   err.Error(),
			Context: map[string]interface{}{
				"err": err,
			},
		})
		log.Fatal("database connection failed: ", err)
	}
	logger.Info(models.InfoLog{
		Message: "Database connected",
		Event:   "database_connected",
	})
	defer db.Close()
	services := services.CreateNewServices(repositories.NewUserRepository(db), repositories.NewProductRepository(db), repositories.NewCycleRepository(db), repositories.NewInstitutionRepository(db), logger)
	handler := handlers.NewHandler(*services)

	hashed, err := utils.HashPassword("123456")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hashed)
	services.UserService.Add("admin", "123456", "admin", "admin", "admin")
	ok := utils.CheckPasswordHash("123456", hashed)
	fmt.Println(ok)

	handler.Init()
}
