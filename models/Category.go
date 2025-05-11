package models

type CategoryData struct {
	Locations  []string `json:"location"`
	Types      []string `json:"type"`
	Statuses   []string `json:"status"`
	Categories []string `json:"category"`
}