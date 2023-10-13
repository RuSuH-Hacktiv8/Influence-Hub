package repository

import (
	"context"
	"influence-hub-influencer/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) UpdateFollowerCount(influencer models.Influencer) error {
	coll := r.DB.Collection("influencer")

	filter := bson.M{"_id": influencer.ID}
	update := bson.M{"$set": bson.M{"instagram_followers": influencer.InstagramFollowers}}
	result := coll.FindOneAndUpdate(context.TODO(), filter, update)
	err := result.Err()
	if err != nil {
		return err
	}
	return nil
}
