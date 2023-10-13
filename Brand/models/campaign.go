package models

type Campaign struct {
	ID               uint   `json:"id,omitempty" gorm:"primaryKey"`
	BrandID          uint    `json:"brand_id,omitempty"`
	Name             string `json:"name,omitempty"`
	Payment          int    `json:"payment,omitempty"`
	MinimumFollowers int    `json:"minimum_followers,omitempty"`
	StartDate        string `json:"start_date,omitempty"`
	EndDate          string `json:"end_date,omitempty"`
	ImageURL         string `json:"image_url,omitempty"`
}
