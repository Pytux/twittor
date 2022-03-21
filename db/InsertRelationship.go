package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"
)

func InsertRelationship(relationship models.Relationship) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := mongoClient.Database("twittor")

	col := db.Collection("relationships")

	_, err := col.InsertOne(ctx, relationship)
	if err != nil {
		return false, err
	}

	return true, nil
}
