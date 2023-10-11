package repository

import (
	"influence-hub-brand/models"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) Repository { return Repository{db} }

func (r Repository) AddBrand(brand models.Brand) (uint, error) {
	result := r.DB.Create(&brand)
	return brand.ID, result.Error
}
