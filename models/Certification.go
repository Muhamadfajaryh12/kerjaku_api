package models

import "time"

type Certification struct {
	ID                int `gorm:"autoIncrement;primaryKey"`
	CertificationName string
	Publisher         string
	EffectiveDate     time.Time
	IDUser	uint
}

type CertificationForm struct {
	CertificationName string `form:"certification_name" json:"certification_name" validate:"required"`
	Publisher         string `form:"publisher" json:"publisher" validate:"required"`
	EffectiveDate     time.Time `form:"effective_date" json:"effective_date" validate:"required"`
}