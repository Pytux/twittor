package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"
)

func DeleteRelationship(relationship models.Relationship) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := mongoClient.Database("twittor")
	col := db.Collection("relationships")

	_, err := col.DeleteOne(ctx, relationship)
	if err != nil {
		return false, err
	}

	return true, nil

}
