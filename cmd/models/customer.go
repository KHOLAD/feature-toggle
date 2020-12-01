package models

// Customer - base customer struct
type Customer struct {
	// CustomerID - base customer id
	CustomerID string `json:"customerId" bson:"customerId"`
	// Features - main features reference
	Features []string `json:"features" bson:"features"`
}
