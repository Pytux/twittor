package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(user models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := mongoClient.Database("twittor")
	col := db.Collection("users")

	profile := make(map[string]interface{})

	if len(user.Name) > 0 {
		profile["name"] = user.Name
	}

	if len(user.Username) > 0 {
		profile["username"] = user.Username
	}

	profile["dateOfBirth"] = user.DateOfBirth

	if len(user.Avatar) > 0 {
		profile["avatar"] = user.Avatar
	}

	if len(user.Banner) > 0 {
		profile["banner"] = user.Banner
	}

	if len(user.Biography) > 0 {
		profile["biography"] = user.Biography
	}

	if len(user.Ubication) > 0 {
		profile["ubication"] = user.Ubication
	}

	if len(user.WebSite) > 0 {
		profile["webSite"] = user.WebSite
	}

	updateString := bson.M{
		"$set": profile,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
