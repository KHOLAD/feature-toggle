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
	e.HideBanner = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Custom Error handler
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// Routes
	e.GET("api/features", handlers.GetFeatures)
	e.POST("api/feature", handlers.CreateFeature)
	e.PUT("api/feature/:id", handlers.UpdateFeature)
	e.GET("api/customers", handlers.GetAllCustomers)
	e.GET("api/customer/:id", handlers.GetCustomerFeatures)
	e.PUT("api/toggle/:customerId/:name", handlers.ToggleFeature)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
