package models

// Error is a global JSON object sent
// between client and server to communicate errors.
type Error struct {
	Context string `json:"context"`
	Origin  string `json:"origin"`
	Params  string `json:"params"`
}
