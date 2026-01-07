package models

import "time"

type Experience struct {
	ID             int  `gorm:"primaryKey;autoIncrement" json:"id"`
	NameCompany		string` json:"name_company" `
	NameExperience string `  json:"name_experience" `
	DateStart      time.Time `json:"date_start" `
	DateEnd		   time.Time ` json:"date_end" `
	Position	   string	 ` json:"position" `
	UserID     uint  ` form:"user_id" json:"user_id"`
}

type ExperienceForm struct {
	NameCompany		string` form:"name_company" json:"name_company" validate:"required"`
	NameExperience string ` form:"name_experience" json:"name_experience" validate:"required"`
	DateStart      time.Time `form:"date_start" json:"date_start" validate:"required"`
	DateEnd		   time.Time ` form:"date_end" json:"date_end" validate:"required"`
	Position	   string	 ` form:"position" json:"position" validate:"required"`
}