package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRelationship(relationship models.Relationship) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := mongoClient.Database("twittor")
	col := db.Collection("relationships")

	condition := bson.M{
		"userid":             relationship.UserID,
		"userrelationshipid": relationship.UserRelationshipID,
	}

	var result models.Relationship

	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}

	return true, nil

}
