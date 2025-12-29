package models

import "time"

type Experience struct {
	ID             int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	NameCompany		string`gorm:"type:varchar(255)" form:"name_company" json:"name_company" validate:"required"`
	NameExperience string `gorm:"type:varchar(255)" form:"name_experience" json:"name_experience" validate:"required"`
	DateStart      time.Time `gorm:"type:date" form:"date_start" json:"date_start" validate:"required"`
	DateEnd		   time.Time `gorm:"type:date" form:"date_end" json:"date_end" validate:"required"`
	Position	   string	 `gorm:"type:varchar(255)" form:"position" json:"position" validate:"required"`
	IDProfile      int64   `gorm:"index" form:"id_profile" json:"id_profile"`
}