package db

import (
	"context"
	"time"

	"github.com/pytux/twittor/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllUsers(ID string, page int64, search string, searchtype string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := mongoClient.Database("twittor")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var user models.User

		err := cur.Decode(&user)
		if err != nil {
			return results, false
		}

		var relationship models.Relationship

		relationship.UserID = ID
		relationship.UserRelationshipID = user.ID.Hex()

		include = false

		found, _ = GetRelationship(relationship)
		if searchtype == "new" && !found {
			include = true
		}

		if searchtype == "follow" && !found {
			include = true
		}

		if relationship.UserRelationshipID == ID {
			include = false
		}

		if include {
			user.Password = ""
			user.Email = ""
			user.WebSite = ""
			user.Banner = ""

			results = append(results, &user)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}

	cur.Close(ctx)

	return results, true
}
