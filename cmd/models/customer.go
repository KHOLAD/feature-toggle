package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Customer - base customer struct
type Customer struct {
	// CustomerID - base customer id
	ID *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	// Features - main features reference
	Features []string `json:"features" bson:"features"`
}
