package models

import "time"

type Certification struct {
	ID                int `gorm:"autoIncrement;primaryKey" json:"id"`
	CertificationName string `json:"certification_name"`
	Publisher         string  `json:"publisher"`
	EffectiveDate     time.Time `json:"effective_date"`
	UserID	uint `json:"user_id"`
}

type CertificationForm struct {
	CertificationName string `form:"certification_name" json:"certification_name" validate:"required"`
	Publisher         string `form:"publisher" json:"publisher" validate:"required"`
	EffectiveDate     time.Time `form:"effective_date" json:"effective_date" validate:"required"`
}