package db

import (
	"database/sql"
	"fmt"
)

// Ensure this is aligned with github.com/raygervais/xavier/shared/models
func (db Database) initializeLogsTable() error {
	query := `
		CREATE VIRTUAL TABLE IF NOT EXISTS logs USING fts4 (
			date 	TEXT NOT NULL,
			type 	TEXT NOT NULL,
			data	TEXT NOT NULL,
		)
	`

	return db.initProvidedTable(query, "logs")
}

// GetAllLogEntries handles querying the DB for all log entries
func (db Database) GetAllLogEntries() (*sql.Rows, error) {
	query := `SELECT rowid, date, type, data FROM logs`
	return db.execQueryStatement(query)
}

// SearchLogsTable handles querying the DB for all log entires
// which comply with the search string
func (db Database) SearchLogsTable(
	search string,
	limit int,
) (*sql.Rows, error) {
	if len(search) == 0 {
		return nil, fmt.Errorf("Invalid search parameter passed in")
	}

	query := fmt.Sprintf(`
		SELECT rowid, date, type, data 
		FROM logs
		WHERE date match ?
		OR type MATCH ?
		OR data MATCH ?
		LIMIT %d
	`, limit)

	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to create search query: %s", err)
	}

	return stmt.Query(search, search, search)
}
