package models

type Vacany struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	NameVacancy string `gorm:"varchar(255)" form:"name_vacany" json:"name_vacany"`
	Description string `gorm:"varchar(255)" form:"description" json:"description"`
	Location    string `gorm:"varchar(255)" form:"location" json:"location"`
	Qty         int64  `gorm:"int" form:"qty" json:"qty"`
	Salary      int64  `gorm:"int" form:"salary" json:"salary"`
}