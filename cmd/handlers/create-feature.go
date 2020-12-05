package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KHOLAD/feature-toggle-api/models"
	"github.com/KHOLAD/feature-toggle-api/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateFeature - POST request creation handler
func CreateFeature(c echo.Context) (err error) {
	f := new(models.Feature)
	if err = c.Bind(f); err != nil {
		return err
	}

	err = f.Validate()
	if err != nil {
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	// Get mongo client
	mongoclient, err := mongo.GetClient()
	if err != nil {
		return mongo.GetClientError()
	}

	// Insert new feature to collection with new object ID
	f.ID = primitive.NewObjectID()
	featCol := mongoclient.Database(mongo.Database).Collection(mongo.FeaturesCollection)
	found, err := featCol.CountDocuments(context.TODO(), bson.M{"technicalName": f.TechnicalName})
	if err != nil {
		em := fmt.Sprintf("Cannot validate document for feature with [%v].", f.TechnicalName)
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}
	if found >= 1 {
		em := fmt.Sprintf("Feature with [%v - %v] already exist.", f.TechnicalName, f.ID.Hex())
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
	}

	_, err = featCol.InsertOne(context.TODO(), f)
	if err != nil {
		em := fmt.Sprintf("Cannot insert feature with [%v].", f.TechnicalName)
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	// Gets customers collection and update with new feature
	cusColl := mongoclient.Database(mongo.Database).Collection(mongo.CustomersCollection)
	for _, cID := range f.CustomerIds {
		findBy := bson.M{"_id": cID}
		change := bson.M{
			"$push": bson.M{"features": models.GetUserEntity(f)},
		}

		_, err = cusColl.UpdateOne(
			context.TODO(),
			findBy,
			change,
		)

		if err != nil {
			em := fmt.Sprintf("Cannot find and insert feature [%v] for customer [%v].", f.TechnicalName, cID.Hex())
			return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
		}
	}

	return c.JSON(http.StatusOK, f)
}
