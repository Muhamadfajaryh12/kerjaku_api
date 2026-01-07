package models

import "time"

type Profile struct {
	ID          int64         `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `json:"name"`
	Summary     string        `json:"summary"`
	Address     string        `json:"address"`
	Phone       int64         `json:"phone"`
	Birth       time.Time     `json:"birth"`
	Email       string       `json:"email"`
	CV          string        `json:"cv"`
	Photo       string      	`json:"photo"`
	Gender 		string 			`json:"gender"`
	UserID      uint          `json:"user_id"`	
}

type ProfileForm struct {
	Name        string         	`json:"name" form:"name" validate:"required"`
	Summary     string       	`json:"summary" form:"summary" validate:"required"`
	Address     string        	`json:"address" form:"address" validate:"required"`
	Phone       int64         	`json:"phone" form:"phone" validate:"required"`
	Birth       time.Time     	`json:"birth" form:"birth" validate:"required"`
	Email       string       	`json:"email" form:"email" validate:"required"`
	CV          string        	`json:"cv" form:"cv" validate:"required"`
	Photo       string      	`json:"photo" form:"photo" validate:"required"`
	Gender 		string 			`json:"gender" form:"gender" validate:"required"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
	Education []Education `json:"education"`
	Language []Language `json:"language"`
	Certification []Certification `json:"certification"`
	Skill []Skill `json:"skill"`
	Experience []Experience `json:"experience"`
}