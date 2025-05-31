package models

import "time"

type Application struct {
	ID          int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Note        string    `gorm:"varchar(255)" form:"note" json:"note"`
	Status      string    `gorm:"varchar(255)" form:"status" json:"status"`
	Date        time.Time `gorm:"type:date" form:"date" json:"date"`
	CoverLetter string    `gorm:"varchar(255)" form:"cover_letter" json:"cover_letter"`
	IDVacancy   int64     `gorm:"index" form:"id_vacancy" json:"id_vacancy"`
	IDProfile   int64     `gorm:"index" form:"id_profile" json:"id_profile"`
	Profile     Profile   `gorm:"foreignKey:IDProfile" json:"profile"`
	Vacancy     Vacancy   `gorm:"foreignKey:IDVacancy" json:"vacancy"`
}

type UpdateApplication struct{
	Status string `form:"status" json:"status"`
	Note   string `form:"note" json:"note"`
}

type DashboardApplication struct {
	TotalData struct{
		TotalApplicant int64 `json:"total_applicant"`
		TotalWaiting   int64 `json:"total_waiting"`
		TotalAssesment	int64 `json:"total_assesment"`
		TotalInterview int64 `json:"total_interview"`
		TotalCompleted int64 `json:"total_completed"`
		TotalRejected int64 `json:"total_rejected"`
	} `json:"total_data"`
	TotalApplicantByName[] struct{
		NameVacancy string `json:"name_vacancy"`
        Count       int64  `json:"count"`
	}`json:"total_applicant_by_name"`
	TotalApplicantByMonth[] struct{
		NameMonth string `json:"name_month"`
		Count int64 `json:"count"`
	}`json:"total_applicant_by_month"`
}