package handlers

import (
	"context"
	"net/http"

	"github.com/KHOLAD/feature-toggle-api/models"
	"github.com/KHOLAD/feature-toggle-api/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

// GetFeatures - Gets all available features
func GetFeatures(c echo.Context) (err error) {
	mongoclient, err := mongo.GetClient()
	if err != nil {
		return mongo.GetClientError()
	}

	featCol := mongoclient.Database(mongo.Database).Collection(mongo.FeaturesCollection)

	cur, err := featCol.Find(context.TODO(), bson.M{})
	if err != nil {
		em := "Cannot find any features"
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	var features []models.Feature
	if err = cur.All(context.TODO(), &features); err != nil {
		em := "Cannot read documents from feature collection"
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	return c.JSON(http.StatusOK, features)
}
