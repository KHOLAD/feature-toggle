package mongo

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionString = "mongodb://admin:password@datastore:27017/featureTable/?authSource=admin"
	// Database - main feature table database
	Database = "featureTable"
	// FeaturesCollection - collection name
	FeaturesCollection = "features"
	// CustomersCollection - collection name
	CustomersCollection = "customers"
)

var mongoInstance *mongo.Client
var connectionError error
var syncOnce sync.Once

// GetClient - Connects to mongo and return client type
func GetClient() (*mongo.Client, error) {
	syncOnce.Do(func() {
		log.Println("Connecting to mongo...")
		clientOptions := options.Client().ApplyURI(connectionString)
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			connectionError = err
		}
		connectionError = client.Ping(context.Background(), nil)

		log.Println("Connected to mongo")
		mongoInstance = client
	})
	return mongoInstance, connectionError
}
