package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Username    string             `bson:"username" json:"username,omitempty"`
	DateOfBirth time.Time          `bson:"dateOfBirth" json:"dateOfBirth,omitempty"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password,omitempty"`
	Avatar      string             `bson:"avatar" json:"avatar,omitempty"`
	Banner      string             `bson:"banner" json:"banner,omitempty"`
	Biography   string             `bson:"biography" json:"biography,omitempty"`
	Ubication   string             `bson:"ubication" json:"ubication,omitempty"`
	WebSite     string             `bson:"webSite" json:"webSite,omitempty"`
}
