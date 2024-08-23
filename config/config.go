package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func Load() {
	// .env dosyasını yükle
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Ortam değişkenlerini al
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	fmt.Println("ortam degiskenleri")
	fmt.Println(DBHost, DBPort, DBUser)
}
