package models

type CategoryData struct {
	Locations  []string `json:"location"`
	Types      []string `json:"type"`
	Categories []string `json:"category"`
}

type CategoryCompanyData struct {
	CompanyType []string `json:"type"`
	Location    []string `json:"location"`
}