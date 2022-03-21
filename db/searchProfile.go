package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchProfile(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*15)
	defer cancel()

	db := mongoClient.Database("twittor")
	col := db.Collection("users")

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": objID}

	err := col.FindOne(ctx, condition).Decode(&profile)

	profile.Password = ""

	if err != nil {
		return profile, err
	}
	return profile, nil
}
