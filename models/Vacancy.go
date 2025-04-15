package models

import (
	"time"
)

type Vacany struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	NameVacancy string    `gorm:"varchar(255)" form:"name_vacancy" json:"name_vacancy"`
	Description string    `gorm:"varchar(255)" form:"description" json:"description"`
	Location    string    `gorm:"varchar(255)" form:"location" json:"location"`
	Qty         int64     `gorm:"int" form:"qty" json:"qty"`
	Salary      int64     `gorm:"int" form:"salary" json:"salary"`
	DateEnd     time.Time `gorm:"type:date" form:"date_end" json:"date_end"`
	DateStart 	time.Time `gorm:"type:date" form:"date_start" json:"date_start"`
	Status		string 	  `gorm:"varchar(255)" form:"status" json:"status"`
	IDCompany 	int64	  `gorm:"index" form:"id_company" json:"id_company"`
	Company 	Company    `gorm:"foreignKey:IDCompany" json:"company"`			
}


