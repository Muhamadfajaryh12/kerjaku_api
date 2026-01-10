package models

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(255)" form:"email" json:"email" validate:"required"`
	Password string `gorm:"type:varchar(255)" form:"password" json:"password" validate:"required"`
	Role     string `form:"role" json:"role" validate:"required"`
}

type LoginForm struct {
	Email    string `form:"email" json:"email" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}