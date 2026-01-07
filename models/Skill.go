package models

type Skill struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Skill  string `json:"skill"`
	UserID uint   `json:"user_id"`
}

type SkillForm struct {
	Skill string `form:"skill" json:"skill" validate:"required"`
}