package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Influencer struct {
	ID                primitive.ObjectID `json:"id,omitempty" gorm:"primaryKey"`
	Email             string             `json:"email,omitempty" gorm:"unique;"`
	Password          string             `json:"password,omitempty"`
	Name              string             `json:"name,omitempty"`
	InstagramUsername string             `json:"instagram_username,omitempty"`
	ImageURL          string             `json:"image_url,omitempty" gorm:"column:image_url"`
}
