package models

type Relationship struct {
	UserID             string `bson:"userid" json:"userId"`
	UserRelationshipID string `bson:"userrelationshipid" json:"userRelationshipId"`
}
