package config

import (
	"influence-hub-brand/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	url := os.Getenv("DB_POSTGRESQL")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&models.Brand{}); err != nil {
		log.Fatal(err)
	}
	return db
}
