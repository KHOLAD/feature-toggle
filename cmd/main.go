package main

import (
	"log"

	"github.com/KHOLAD/feature-toggle-api/handlers"
	"github.com/KHOLAD/feature-toggle-api/mongo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// MongoDB instance
	_, dbConnectionError := mongo.GetClient()
	if dbConnectionError != nil {
		log.Fatal(dbConnectionError)
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/feature", handlers.CreateFeature)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
