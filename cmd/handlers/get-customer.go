package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KHOLAD/feature-toggle-api/models"
	"github.com/KHOLAD/feature-toggle-api/mongo"
	m "github.com/KHOLAD/feature-toggle-api/mongo"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetCustomerFeatures - get customer features by customer id
func GetCustomerFeatures(c echo.Context) (err error) {
	mongoclient, err := m.GetClient()
	if err != nil {
		return m.GetClientError()
	}

	id := c.Param("id")
	cusID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		em := fmt.Sprintf("Cannot parse customer ID from [%v]", id)
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
	}

	cusColl := mongoclient.Database(m.Database).Collection(m.CustomersCollection)
	uf := models.CustomerFeatures{}
	findBy := bson.M{"_id": cusID}
	err = cusColl.FindOne(context.TODO(), findBy).Decode(&uf)
	if err != nil {
		em := fmt.Sprintf("Cannot find customer with [%v]", id)
		return models.NewHTTPError(http.StatusBadRequest, "BadRequest", em)
	}

	return c.JSON(http.StatusOK, uf)
}

// GetAllCustomers - list of all available customers
func GetAllCustomers(c echo.Context) (err error) {
	mongoclient, err := m.GetClient()
	if err != nil {
		return m.GetClientError()
	}

	cusColl := mongoclient.Database(mongo.Database).Collection(mongo.CustomersCollection)

	cur, err := cusColl.Find(context.TODO(), bson.M{})
	if err != nil {
		em := "Cannot find any customers"
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	var customers []models.AvailableCustomers
	if err = cur.All(context.TODO(), &customers); err != nil {
		em := "Cannot read documents from customers collection"
		return models.NewHTTPError(http.StatusInternalServerError, "InternalServerError", em)
	}

	return c.JSON(http.StatusOK, customers)
}
