package main

import (
	"log"
	"net/http"

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
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
