package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KHOLAD/feature-toggle-api/models"
	m "github.com/KHOLAD/feature-toggle-api/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func containsID(s []primitive.ObjectID, e primitive.ObjectID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// UpdateFeature - update certain feature
func UpdateFeature(c echo.Context) (err error) {
	f := new(models.Feature)
	if err = c.Bind(f); err != nil {
		return err
	}

	err = f.Validate()
	if err != nil {
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	id := c.Param("id")
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		em := fmt.Sprintf("Cannot parse ID from [%v]", id)
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
	}

	mongoclient, err := m.GetClient()
	if err != nil {
		em := "Cannot get client from Database"
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	f.ID = docID
	featCol := mongoclient.Database(m.Database).Collection(m.FeaturesCollection)

	findBy := bson.M{"_id": docID}
	change := bson.M{
		"$set": bson.M{
			"customerIds":   f.CustomerIds,
			"inverted":      f.Inverted,
			"displayName":   f.DisplayName,
			"technicalName": f.TechnicalName,
			"expiresOn":     f.ExpiresOn,
			"description":   f.Description,
		},
	}

	_, err = featCol.UpdateOne(
		context.TODO(),
		findBy,
		change,
	)

	if err != nil {
		fmt.Println(err)
		em := fmt.Sprintf("Cannot update feature [%v - %v].", f.TechnicalName, id)
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
	}

	return c.JSON(http.StatusOK, f)
}
