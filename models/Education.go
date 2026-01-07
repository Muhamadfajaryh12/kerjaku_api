package models

import "time"

type Education struct {
	ID            int `gorm:"primaryKey;autoIncrement" json:"id"`
	EducationName string `json:"education_name"`
	Major         string `json:"major"`
	GraduateDate  time.Time `json:"graduate_date"`
	Level string `json:"level"`
	UserID uint `json:"user_id"`
}

type EducationForm struct {
	EducationName string `form:"education_name" json:"education_name" validate:"required"`
	Major string `form:"major" json:"major" validate:"required"`
	GraduateDate time.Time `form:"graduate_date" json:"graduate_date" validate:"required"`
	Level string `form:"level" json:"level" validate:"required"`
}