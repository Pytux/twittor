package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnFollowedTweets struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"userid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"userrelationshipid" json:"userRelationshipId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"message" json:"message,omitempty"`
		Fecha   time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
