package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFollowedTweets(ID string, page int) ([]models.ReturnFollowedTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := mongoClient.Database("twittor")
	col := db.Collection("relationships")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)

	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userrelationshipid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cur, _ := col.Aggregate(ctx, conditions)

	var result []models.ReturnFollowedTweets

	err := cur.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
