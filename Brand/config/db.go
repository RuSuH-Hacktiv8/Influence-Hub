package config

import (
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

	return db
}
