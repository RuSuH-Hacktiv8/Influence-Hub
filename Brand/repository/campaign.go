package repository

import (
	"errors"
	"influence-hub-brand/models"

	"gorm.io/gorm"
)

func (r *Repository) AddCampaign(campaign models.Campaign) (models.Campaign, error) {
	query := r.DB.Create(&campaign)
	if query.Error != nil {
		return models.Campaign{}, query.Error
	}

	return campaign, nil
}

func (r *Repository) GetCampaign(brandid int) ([]models.Campaign, error) {
	campaign := []models.Campaign{}
	query := r.DB.Table("campaigns").Where("BrandID=?", brandid).Find(campaign)
	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return []models.Campaign{}, errors.New("user not found")
		}

		return []models.Campaign{}, query.Error
	}

	return campaign, nil
}

func (r *Repository) GetAllCampaign() ([]models.Campaign, error) {
	campaigns := []models.Campaign{}
	query := r.DB.Table("campaigns").Find(&campaigns)
	if query.Error != nil {
		return []models.Campaign{}, query.Error
	}

	return campaigns, nil
}

func (r *Repository) EditCampaign(id uint, updatedCampaign models.Campaign) (models.Campaign, error) {
	var existingCampaign models.Campaign

	if err := r.DB.First(&existingCampaign, id).Error; err != nil {
		return models.Campaign{}, err
	}

	r.DB.Model(&existingCampaign).Updates(updatedCampaign)

	return existingCampaign, nil
}

func (r *Repository) DeletesCampaign(id uint) error {
	var existingCampaign models.Campaign

	if err := r.DB.First(&existingCampaign, id).Error; err != nil {
		return err
	}

	if err := r.DB.Delete(&existingCampaign).Error; err != nil {
		return err
	}

	return nil
}
