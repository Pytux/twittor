package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient = connectDB()

var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/twittor")

func connectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Successful DB connection")

	return client

}

func CheckConnection() bool {
	err := mongoClient.Ping(context.TODO(), nil)

	if err != nil {
		return false
	}

	return true
}
