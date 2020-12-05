package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer - base customer struct
type Customer struct {
	// CustomerID - base customer id
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	// Features - main features reference
	Features []UserFeature `json:"features" bson:"features"`
}

// CustomerFeatures - features per customers type
type CustomerFeatures struct {
	// CustomerID - base customer id
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Features []UserFeature      `json:"features" bson:"features"`
}

// AvailableCustomers - list of available customers type
type AvailableCustomers struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
