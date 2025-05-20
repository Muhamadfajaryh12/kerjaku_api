package models

import "time"

type Profile struct {
	ID          int64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string        `gorm:"type:varchar(255)" form:"name" json:"name" validate:"required"`
	Summary     string        `gorm:"type:varchar(255)" form:"summary" json:"summary" validate:"required"`
	Education   string        `gorm:"type:varchar(255)" form:"education" json:"education" validate:"required"`
	Address     string        `gorm:"type:varchar(255)" form:"address" json:"address" validate:"required"`
	Phone       int64         `gorm:"type:bigint" form:"phone" json:"phone" validate:"required"`
	Birth       time.Time     `gorm:"type:date" form:"birth" json:"birth" validate:"required"`
	Email       string        `gorm:"type:varchar(255)" form:"email" json:"email" validate:"required"`
	Skills      []string      `gorm:"type:text;serializer:json" form:"skills" json:"skills" validate:"required"`
	CV          string        `gorm:"type:varchar(255)" form:"cv" json:"cv" validate:"required"`
	Photo       string        `gorm:"type:varchar(255)" form:"photo" json:"photo" validate:"required"`
	IDUser      int64         `gorm:"type:bigint" json:"id_user" form:"id_user" validate:"required"`
	Experience  []Experience  `gorm:"foreignKey:IDProfile;" json:"experience"`
	Application []Application `gorm:"foreignKey:IDProfile" json:"application"`
}