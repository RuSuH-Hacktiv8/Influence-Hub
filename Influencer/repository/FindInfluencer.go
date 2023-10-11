package repository

import (
	"context"
	"fmt"
	"influence-hub-influencer/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindByEmail(email string) (models.Influencer, error) {
	collection := r.DB.Collection("influencer")

	var influencer models.Influencer

	filter := bson.M{"email": email}
	if err := collection.FindOne(context.Background(), filter).Decode(&influencer); err != nil {
		return models.Influencer{}, fmt.Errorf(err.Error())
	}

	return influencer, nil
}

func (r *Repository) FindById(id string) (models.Influencer, error) {
	collection := r.DB.Collection("influencer")

	var influencer models.Influencer

	filter := bson.M{"_id": id}
	if err := collection.FindOne(context.Background(), filter).Decode(&influencer); err != nil {
		return models.Influencer{}, fmt.Errorf(err.Error())
	}

	return influencer, nil
}
