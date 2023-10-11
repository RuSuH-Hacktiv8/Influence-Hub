package repository

import (
	"context"
	"influence-hub-influencer/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *Repository) AddInfluencer(influencer models.Influencer) (primitive.ObjectID, error) {
	collection := r.DB.Collection("influencer")

	result, err := collection.InsertOne(context.TODO(), influencer)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}
