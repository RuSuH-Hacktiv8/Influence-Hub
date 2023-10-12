package models

type Campaign struct {
	ID               uint   `json:"id,omitempty" bson:"_id,omitempty"`
	BrandID          int    `json:"brand_id,omitempty" bson:"brand_id,omitempty"`
	Name             string `json:"name,omitempty" bson:"name,omitempty"`
	Payment          int    `json:"payment,omitempty" bson:"payment,omitempty"`
	MinimumFollowers int    `json:"minimum_followers,omitempty" bson:"minimum_followers,omitempty"`
	StartDate        string `json:"start_date,omitempty" bson:"start_date,omitempty"`
	EndDate          string `json:"end_date,omitempty" bson:"end_date,omitempty"`
	ImageURL         string `json:"image_url,omitempty" bson:"image_url,omitempty"`
}
