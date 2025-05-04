package models

type Application struct {
	ID        int64   `gorm:"primaryKey;autoIncrement" json:"id"`
	Note      string  `gorm:"varchar(255)" form:"note" json:"note"`
	Status    string  `gorm:"varchar(255)" form:"status" json:"status"`
	IDVacancy int64   `gorm:"index" form:"id_vacancy" json:"id_vacancy"`
	IDUser    int64   `gorm:"index" form:"id_user" json:"id_user"`
	Profile   Profile `gorm:"foreignKey:IDUser" json:"user"`
	Vacancy   Vacancy `gorm:"foreignKey:IDVacancy" json:"vacancy"`
}