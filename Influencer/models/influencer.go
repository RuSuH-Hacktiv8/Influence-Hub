package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Influencer struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email              string             `json:"email,omitempty" bson:"email,omitempty"`
	Password           string             `json:"password,omitempty" bson:"password,omitempty"`
	Name               string             `json:"name,omitempty" bson:"name,omitempty"`
	InstagramUsername  string             `json:"instagram_username,omitempty" bson:"instagram_username,omitempty"`
	InstagramFollowers int                `json:"instagram_followers,omitempty" bson:"instagram_followers,omitempty"`
	ImageURL           string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
}
