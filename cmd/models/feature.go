package models

import (
	"sync"
	"time"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate *validator.Validate
var syncOnce sync.Once

// Feature - main feature entity type
type Feature struct {
	ID            primitive.ObjectID   `bson:"_id" json:"_id,omitempty"`
	DisplayName   string               `json:"displayName" bson:"displayName" validate:"required,min=3"`
	TechnicalName string               `json:"technicalName" bson:"technicalName" validate:"required,min=3"`
	ExpiresOn     time.Time            `json:"expiresOn,omitempty" bson:"expiresOn"`
	Description   string               `json:"description" bson:"description" validate:"max=10"`
	Inverted      bool                 `json:"inverted" bson:"inverted"`
	CustomerIds   []primitive.ObjectID `json:"customerIds,omitempty" validate:"required,min=1"`
}

// UserFeature - user feature entity type
type UserFeature struct {
	Name     string `json:"name" bson:"name"`
	Active   bool   `json:"active" bson:"active"`
	Inverted bool   `json:"inverted" bson:"inverted"`
	Expired  bool   `json:"expired" bson:"expired"`
}

// GetUserEntity - returns user entity from Feature type
func GetUserEntity(f *Feature) UserFeature {
	return UserFeature{
		Name:     f.TechnicalName,
		Active:   false,
		Inverted: f.Inverted,
		Expired:  false,
	}
}

// Validate - Feature validation
func (f *Feature) Validate() (err error) {
	syncOnce.Do(func() {
		validate = validator.New()
	})

	err = validate.Struct(f)
	if err != nil {
		return err
	}

	return nil
}
