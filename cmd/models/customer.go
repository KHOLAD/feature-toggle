package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer - base customer struct
type Customer struct {
	// CustomerID - base customer id
	ID *primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// Features - main features reference
	Features []UserFeature `json:"features" bson:"features"`
}
