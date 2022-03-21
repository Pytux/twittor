package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"
	"github.com/pytux/twittor/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveUser(user models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := mongoClient.Database("twittor")

	col := db.Collection("users")

	user.Password, _ = utils.EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	Id, _ := result.InsertedID.(primitive.ObjectID)

	return Id.String(), true, nil
}
