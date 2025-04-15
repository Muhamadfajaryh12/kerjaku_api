package models

type Company struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id" `
	CompanyName string `gorm:"type:varchar(255)" form:"company_name" json:"company_name"`
	CompanyType string `gorm:"type:varchar(255)" form:"company_type" json:"company_type"`
	Location    string `gorm:"type:varchar(255)" form:"location" json:"location"`
	Size        int64  `gorm:"type:int" form:"size" json:"size"`
	Photo       string `gorm:"type:varchar(255)" form:"photo" json:"photo"`
	Description string `gorm:"type:varchar(255)" form:"description" json:"description"`
	IDUser      int64  `gorm:"index" json:"id_user" form:"id_user"`
}