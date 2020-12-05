package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/KHOLAD/feature-toggle-api/models"
	m "github.com/KHOLAD/feature-toggle-api/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ToggleFeature - feature toggle for specific customer
func ToggleFeature(c echo.Context) (err error) {
	customerID := c.Param("customerId")
	cID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		em := fmt.Sprintf("Cannot parse ID from [%v]", customerID)
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
	}
	featureName := c.Param("name")
	mongoclient, err := m.GetClient()
	if err != nil {
		return m.GetClientError()
	}

	cusColl := mongoclient.Database(m.Database).Collection(m.CustomersCollection)

	matchStage := bson.D{primitive.E{Key: "$match", Value: bson.M{"_id": cID}}}
	unwindStage := bson.D{primitive.E{Key: "$unwind", Value: "$features"}}
	featureMatchStage := bson.D{primitive.E{Key: "$match", Value: bson.M{"features.name": featureName}}}
	projectStage := bson.D{primitive.E{Key: "$project", Value: bson.M{
		"name":   "$features.name",
		"active": "$features.active",
	}}}
	opts := options.Aggregate().SetMaxTime(2 * time.Second)

	cursor, err := cusColl.Aggregate(context.TODO(), mongo.Pipeline{matchStage, unwindStage, featureMatchStage, projectStage}, opts)
	if err != nil {
		em := fmt.Sprintf("Cannot aggregate feature [%v] for [%v].", featureName, cID.Hex())
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	// User feature toggle slice
	var cf []models.UserFeature
	if err = cursor.All(context.TODO(), &cf); err != nil {
		em := fmt.Sprintf("Cannot parse cursor, feature [%v] for [%v].", featureName, cID.Hex())
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	// Toggle
	findBy := bson.M{"_id": cID, "features.name": featureName}
	change := bson.M{"$set": bson.M{"features.$.active": !cf[0].Active}}
	_, err = cusColl.UpdateOne(
		context.TODO(),
		findBy,
		change,
	)

	if err != nil {
		em := fmt.Sprintf("Cannot toggle feature [%v] for customer [%v].", featureName, cID.Hex())
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
	}

	cf[0].Active = !cf[0].Active
	return c.JSON(http.StatusOK, cf[0])
}
