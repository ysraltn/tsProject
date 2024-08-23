package database

import (
	"fmt"
	"tsProject/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func Init() (*sqlx.DB, error) {
	// Veritabanı bağlantı bilgilerini config paketinden al
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	// Veritabanı bağlantısını başlat
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// Veritabanına bağlantının başarılı olup olmadığını test et
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Veritabanına başarıyla bağlanıldı!")
	return db, nil
}
