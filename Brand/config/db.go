package config

import (
	"influence-hub-brand/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	url := "postgres://postgres.sjjxkceqhpwrokzycxvp:oq29op4sg7XoTDsg@aws-0-ap-southeast-1.pooler.supabase.com:6543/postgres"
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&models.Brand{}); err != nil {
		log.Fatal(err)
	}
	return db
}
