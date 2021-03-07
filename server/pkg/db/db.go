package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" //Used as SQL driver
)

// Database wraps around *sql.DB connection, imported by Server.
type Database struct {
	connection *sql.DB
}

// CreateDatabaseConnection creates or connects to an existing SQLite3 DB,
// path provided must be absolute path to file.
func CreateDatabaseConnection(path string) (Database, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return Database{}, fmt.Errorf(
			"error opening database connection: \n%s",
			err,
		)
	}

	return Database{
		connection: db,
	}, nil
}

// InitializeTables allows us to create application tables
func (db Database) InitializeTables() error {
	tables := []struct {
		Name string
		Func func() error
	}{
		{
			"version",
			db.initializeApplicationVersion,
		}, {
			"logs",
			db.initializeLogsTable,
		},
	}

	for _, table := range tables {
		err := table.Func()
		if err != nil {
			return err
		}
	}

	return nil
}

// Helper Functions
func (db Database) prepareQueryStatement(query string) (*sql.Stmt, error) {
	return db.connection.Prepare(query)
}

func (db Database) execQueryStatement(query string) (*sql.Rows, error) {
	return db.connection.Query(query)
}

func (db Database) initProvidedTable(query, table string) error {
	stmt, err := db.prepareQueryStatement(query)
	if err != nil {
		return fmt.Errorf("Failed to create logs table: %s", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("Failed to create table %s: %s", table, err)
	}

	return nil
}

// Table Initializers
// - version
// - logs

func (db Database) initializeApplicationVersion() error {
	query := `
		CREATE VIRTUAL TABLE IF NOT EXISTS version using fts4 (
			major NUMBER NOT NULL
			minor NUMBER NOT NULL
			build NUMBER NOT NULL
		)
	`

	return db.initProvidedTable(query, "version")
}

// Ensure this is aligned with github.com/raygervais/xavier/shared/models
func (db Database) initializeLogsTable() error {
	query := `
		CREATE VIRTUAL TABLE IF NOT EXISTS logs USING fts4 (
			date 	TEXT NOT NULL,
			type 	TEXT NOT NULL,
			data	TEXT NOT NULL,
		)`

	return db.initProvidedTable(query, "logs")
}

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
