package models

type Profile struct {
	ID        int64  `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"type:varchar(255)" form:"name" json:"name"`
	Summary   string `gorm:"type:varchar(255)" form:"summary" json:"summary"`
	Education string `gorm:"type:varchar(255)" form:"education" json:"education"`
	Address   string `gorm:"type:varchar(255)" form:"address" json:"address"`
	Skills    string `gorm:"type:nvarchar(MAX)" form:"skills" json:"skills"` // Simpan sebagai string JSON
	IDUser    int64  `gorm:"index" json:"id_user" form:"id_user"`
}