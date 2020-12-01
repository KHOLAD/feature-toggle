package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Feature - main feature entity type
type Feature struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	DisplayName   string             `json:"displayName" bson:"displayName"`
	TechnicalName string             `json:"technicalName" bson:"technicalName"`
	ExpiresOn     time.Time          `json:"expiresOn,omitempty" bson:"expiresOn"`
	Description   string             `json:"description" bson:"description"`
	Inverted      bool               `json:"inverted" bson:"inverted"`
	CustomerIds   []string           `json:"customerIds"`
}

// UserFeature - user feature entity type
type UserFeature struct {
	Name     string `json:"name" bson:"name"`
	Active   bool   `json:"active" bson:"active"`
	Inverted bool   `json:"inverted" bson:"inverted"`
	Expired  bool   `json:"expired" bson:"expired"`
}
