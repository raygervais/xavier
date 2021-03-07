package models

// LogEntry is an example of a retrievable data item
// Each retrievable data item requires entries in:
// - github.com/raygervais/xavier/server/pkg/db
// - github.com/raygervais/xavier/server/pkg/api
type LogEntry struct {
	RowID int    `json:"rowId"`
	Date  string `json:"date"`
	Data  string `json:"data"`
	Type  string `json:"type"`
}
