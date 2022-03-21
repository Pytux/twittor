package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweets(ID string, page int64) ([]*models.GetTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := mongoClient.Database("twittor")
	col := db.Collection("tweets")

	var results []*models.GetTweets

	condition := bson.M{"userid": ID}

	options := options.Find()
	options.SetSkip((page - 1) * 20)
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "date", Value: -1}})

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.GetTweets

		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}

		results = append(results, &register)
	}

	return results, true
}
