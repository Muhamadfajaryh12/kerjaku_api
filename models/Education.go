package models

import "time"

type Education struct {
	ID            int `gorm:"primaryKey;autoIncrement"`
	EducationName string
	Major         string
	GraduateDate  time.Time
	IDUser uint
}

type EducationForm struct {
	EducationName string `form:"education_name" json:"education_name" validate:"required"`
	Major string `form:"major" json:"major" validate:"required"`
	GraduateDate time.Time `form:"graduate_date" json:"graduate_date" validate:"required"`
}