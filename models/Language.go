package models

type Language struct {
	ID       int `gorm:"primaryKey;autoIncrement"`
	Language string
	Level    string
	IDuser   uint
}

type LanguageForm struct {
	Language string `form:"language" json:"language" validate:"required"`
	Level    string `form:"language" json:"language" validate:"required"`
}