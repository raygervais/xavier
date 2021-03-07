package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/raygervais/xavier/server/pkg/api"
	"github.com/raygervais/xavier/server/pkg/db"
)

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "America/New_York",
	}))

	db, err := db.CreateDatabaseConnection("/tmp/db.sqlite")
	if err != nil {
		fmt.Errorf("Failed to create database connection: %s", err)
	}

	err = db.InitializeTables()
	if err != nil {
		fmt.Errorf("Failed to create application tables: %s", err)
	}

	// Application Routes & Coupling to DB
	_ = api.InitializeAPI(app, db)

	app.Listen(":8080")
}
