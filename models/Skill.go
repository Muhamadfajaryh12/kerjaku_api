package models

type Skill struct {
	ID     int `gorm:"primaryKey;autoIncrement"`
	Skill  string
	IDuser uint
}

type SkillForm struct {
	Skill string `form:"skill" json:"skill" validate:"required"`
}