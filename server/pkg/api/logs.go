package api

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/raygervais/xavier/pkg/models"
)

/*
 * API ROUTES
 * URL Entry: /logs/
 */
func (api ApplicationInterface) getAllLogEntries(c *fiber.Ctx) error {
	rows, err := api.database.GetAllLogEntries()
	if err != nil {
		api.errorHandler(c, "getAllLogEntires", err.Error(), "", 400)
		return err
	}

	logs := []models.LogEntry{}
	api.iterateOverLogRows(rows, &logs)

	return c.JSON(logs)

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
	api.iterateOverLogRows(rows, &logs)

	return c.JSON(logs)
}

func (api ApplicationInterface) iterateOverLogRows(rows *sql.Rows, logs *[]models.LogEntry) {
	for rows.Next() {
		log := models.LogEntry{}
		rows.Scan(&log.RowID, &log.Data, &log.Date, &log.Type)
		*logs = append(*logs, log)
	}
}
