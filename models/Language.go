package models

type Language struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Language string `json:"language"`
	Level    string `json:"level"`
	UserID   uint   `json:"user_id"`
}

type LanguageForm struct {
	Language string `form:"language" json:"language" validate:"required"`
	Level    string `form:"level" json:"level" validate:"required"`
}