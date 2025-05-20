package models

import "time"

type Application struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Note        string    `gorm:"varchar(255)" form:"note" json:"note"`
	Status      string    `gorm:"varchar(255)" form:"status" json:"status"`
	Date        time.Time `gorm:"type:date" form:"date" json:"date"`
	CoverLetter string    `gorm:"varchar(255)" form:"cover_letter" json:"cover_letter"`
	IDVacancy   int64     `gorm:"index" form:"id_vacancy" json:"id_vacancy"`
	IDProfile   int64     `gorm:"index" form:"id_profile" json:"id_profile"`
	Profile     Profile   `gorm:"foreignKey:IDProfile" json:"profile"`
	Vacancy     Vacancy   `gorm:"foreignKey:IDVacancy" json:"vacancy"`
}