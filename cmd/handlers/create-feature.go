package handlers

import (
	"context"
	"log"
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
	f.ID = primitive.NewObjectID()

	if err = c.Bind(f); err != nil {
		return
	}
	// Get mongo client
	mongoclient, _ := mongo.GetClient()
	// Insert new feature to collection
	featureCollection := mongoclient.Database(mongo.Database).Collection(mongo.FeaturesCollection)
	_, err = featureCollection.InsertOne(context.TODO(), f)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Gets customers collection and update with new feature
	customersCollection := mongoclient.Database(mongo.Database).Collection(mongo.CustomersCollection)
	for _, cID := range f.CustomerIds {
		userEntity := models.UserFeature{
			Name:     f.TechnicalName,
			Active:   false,
			Inverted: f.Inverted,
			Expired:  false,
		}

		findBy := bson.M{"customerId": cID}
		change := bson.M{"$push": bson.M{"features": userEntity}}

		res := customersCollection.FindOneAndUpdate(
			context.TODO(),
			findBy,
			change,
		)

		if res.Err() != nil {
			log.Fatal(res.Err())
			return res.Err()
		}
	}

	return c.JSON(http.StatusOK, f)
}
