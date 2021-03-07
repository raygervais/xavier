package models

// Version of an artifact, including application.
type Version struct {
	Major int `json:"major"`
	Minor int `json:"minor"`
	Build int `json:"build"`
}
