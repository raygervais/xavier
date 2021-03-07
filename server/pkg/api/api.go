package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raygervais/xavier/pkg/models"
	"github.com/raygervais/xavier/server/pkg/db"
)

// ApplicationInterface is a singleton instance of our API
type ApplicationInterface struct {
	database db.Database
}

// InitializeAPI allows us to configure and
// inject application logic into the API
func InitializeAPI(app *fiber.App, db db.Database) ApplicationInterface {
	api := ApplicationInterface{
		database: db,
	}

	app.Get("/", api.entry)
	app.Get("/logs/", api.getAllLogEntries)
	app.Get("/logs/:search", api.getAllLogsBySearch)

	return api
}

func (api ApplicationInterface) entry(c *fiber.Ctx) error {
	return c.SendString("Xavier, Version 0.1.0!")
}

// wrapper around Fiber's c.Status and c.JSON to reduce code repetition.
func (api ApplicationInterface) errorHandler(
	c *fiber.Ctx,
	origin, context, params string,
	status int,
) {
	c.Status(status)
	c.JSON(models.Error{
		Origin:  origin,
		Context: context,
		Params:  params,
	})
}
