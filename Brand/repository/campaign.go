package repository

import (
	"errors"
	"influence-hub-brand/models"

	"gorm.io/gorm"
)

// type Campaign struct {
// 	ID               uint   `json:"id,omitempty" gorm:"primaryKey"`
// 	BrandID          int    `json:"brand_id,omitempty"`
// 	Name             string `json:"name,omitempty"`
// 	Payment          int    `json:"payment,omitempty"`
// 	MinimumFollowers int    `json:"minimum_followers,omitempty"`
// 	StartDate        string `json:"start_date,omitempty"`
// 	EndDate          string `json:"end_date,omitempty"`
// 	ImageURL         string `json:"image_url,omitempty"`
// }

func (r *Repository) AddCampaign(campaign models.Campaign) (models.Campaign, error) {
	newCampaign := models.Campaign{
		Name:             campaign.Name,
		BrandID:          campaign.BrandID,
		Payment:          campaign.Payment,
		MinimumFollowers: campaign.MinimumFollowers,
		StartDate:        campaign.StartDate,
		EndDate:          campaign.EndDate,
		ImageURL:         campaign.ImageURL,
	}
	query := r.DB.Create(&campaign)
	if query.Error != nil {
		return models.Campaign{}, query.Error
	}

	return newCampaign, query.Error
}

func (r *Repository) GetCampaign(brandid int) ([]models.Campaign, error) {
	campaign := []models.Campaign{}
	query := r.DB.Table("campaigns").Where("BrandID=?", brandid).Find(campaign)
	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return []models.Campaign{}, errors.New("User not found")
		}

		return []models.Campaign{}, query.Error
	}

	return campaign, nil
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
