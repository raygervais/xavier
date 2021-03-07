package models

// LogEntry is an example of a retrievable data item
type LogEntry struct {
	RowID int    `json:"rowId"`
	Date  string `json:"date"`
	Data  string `json:"data"`
	Type  string `json:"type"`
}
