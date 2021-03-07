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

/*
 * API ROUTES
 * URL Entry: /logs
 */
func (api ApplicationInterface) getAllLogEntries(c *fiber.Ctx) error {
	return nil
}

func (api ApplicationInterface) getAllLogsBySearch(c *fiber.Ctx) error {
	searchString := c.Params("search")
	rows, err := api.database.SearchLogsTable(searchString, 10)
	if err != nil {
		api.errorHandler(
			c,
			"getAllLogsBySearch",
			err.Error(),
			searchString,
			400,
		)
		return err
	}

	logs := []models.LogEntry{}

	for rows.Next() {
		log := models.LogEntry{}
		rows.Scan(&log.RowID, &log.Data, &log.Date, &log.Type)
		logs = append(logs, log)
	}

	return c.JSON(logs)
}

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
