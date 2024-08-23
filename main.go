package main

import (
	"log"
	"tsProject/config"
	"tsProject/database"
	_ "tsProject/docs"
	"tsProject/handlers"
	"tsProject/repositories"
	"tsProject/services"
)

// @title Fiber API Documentation
// @version 1.0
// @description This is a sample API with JWT authentication.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.Load()
	db, err := database.Init()
	if err != nil {
		log.Fatal("database connection failed: ", err)
	}
	defer db.Close()
	services := services.CreateNewServices(repositories.NewUserRepository(db), repositories.NewProductRepository(db), repositories.NewCycleRepository(db), repositories.NewInstitutionRepository(db))
	handler := handlers.NewHandler(*services)
	handler.Init()
}
