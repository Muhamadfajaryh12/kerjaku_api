package models

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"type:varchar(255)" form:"username" json:"username" validate:"required"`
	Password string `gorm:"type:varchar(255)" form:"password" json:"password" validate:"required"`
}
