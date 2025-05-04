package models

import (
	"time"
)

type Vacancy struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	NameVacancy string    `gorm:"varchar(255)" form:"name_vacancy" json:"name_vacancy" validate:"required"`
	Description string    `gorm:"varchar(255)" form:"description" json:"description" validate:"required"`
	Location    string    `gorm:"varchar(255)" form:"location" json:"location" validate:"required"`
	Qty         int64     `gorm:"int" form:"qty" json:"qty" validate:"required"`
	Salary      int64     `gorm:"int" form:"salary" json:"salary" validate:"required"`
	DateEnd     time.Time `gorm:"type:date" form:"date_end" json:"date_end" validate:"required"`
	DateStart 	time.Time `gorm:"type:date" form:"date_start" json:"date_start" validate:"required"`
	Status		string 	  `gorm:"varchar(255)" form:"status" json:"status"`
	IDCompany 	int64	  `gorm:"index" form:"id_company" json:"id_company"`
	// Company 	Company    `gorm:"foreignKey:IDCompany" json:"company"`			
}


type IVacancy struct {
    ID          int64     `json:"id"`
    NameVacancy string    `json:"name_vacancy"`
    Description string    `json:"description"`
    Location    string    `json:"location"`
    Qty         int64     `json:"qty"`
    Salary      int64     `json:"salary"`
    DateEnd     time.Time `json:"date_end"`
    DateStart   time.Time `json:"date_start"`
    Status      string    `json:"status"`
    IDCompany   int64     `json:"id_company"`
	Company 	Company	  `json:"company"`
}