package models

type Company struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id" `
	CompanyName string `gorm:"type:varchar(255)" form:"company_name" json:"company_name" validate:"required"`
	CompanyType string `gorm:"type:varchar(255)" form:"company_type" json:"company_type" validate:"required"`
	Location    string `gorm:"type:varchar(255)" form:"location" json:"location" validate:"required"`
	Size        int64  `gorm:"type:int" form:"size" json:"size" validate:"required"`
	Photo       string `gorm:"type:varchar(255)" form:"photo" json:"photo" validate:"required"`
	Description string `gorm:"type:varchar(255)" form:"description" json:"description" validate:"required"`
	IDUser      int64  `gorm:"index" json:"id_user" form:"id_user" validate:"required"`
}

type ICompanyVacancy struct {
	ID          int64     `json:"id" `
	CompanyName string    `json:"company_name"`
	CompanyType string    `json:"company_type"`
	Location    string    `json:"location"`
	Size        int64     `json:"size"`
	Photo       string    `json:"photo"`
	Description string    `json:"description"`
	Vacancy     []Vacancy `json:"vacancy"`
}

type UpdateCompany struct {
	ID          int64  `json:"id" `
	CompanyName string `json:"company_name" form:"company_name"`
	CompanyType string `json:"company_type" form:"company_type"`
	Location    string `json:"location" form:"location"`
	Size        int64  `json:"size" form:"size"`
	Photo       string `json:"photo" form:"photo"`
	Description string `json:"description" form:"description"`
}

type CompanyFilter struct {
	Search   string `query:"search"`
	Location string `query:"location"`
	Type     string `query:"type"`
}