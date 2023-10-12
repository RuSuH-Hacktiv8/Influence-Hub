package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	url := os.Getenv("DB_POSTGRESQL")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: url,

		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
