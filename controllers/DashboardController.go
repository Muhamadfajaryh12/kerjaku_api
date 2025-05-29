package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func DashboardApplication(c *fiber.Ctx) error {
	var application models.DashboardApplication
	var err error

	err = databases.DB.Table("applications").
		Select(`
			COUNT(CASE WHEN status = 'waiting' THEN 1 END) as total_waiting,
			COUNT(CASE WHEN status = 'assessment' THEN 1 END) as total_assesment,
			COUNT(CASE WHEN status = 'interview' THEN 1 END) as total_interview,
			COUNT(CASE WHEN status = 'completed' THEN 1 END) as total_completed,
			COUNT(CASE WHEN status = 'rejected' THEN 1 END) as total_rejected,
			COUNT(*) as total_applicant
		`).Scan(&application.TotalData).Error

	err = databases.DB.Table("applications").
		Select(`
			vacancies.name_vacancy,
			COUNT(applications.id) AS count
		`).
		Joins(" JOIN vacancies ON vacancies.id = applications.id_vacancy").
		Group("vacancies.name_vacancy").
		Scan(&application.TotalApplicantByName).Error

	err = databases.DB.Table("applications").
		Select(`
			DATENAME(MONTH,date) AS name_month,
			COUNT (id) AS count
		`).
		Group("DATENAME(MONTH,date)").
		Scan(&application.TotalApplicantByMonth).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to fetch dashboard data",
			"error":   err.Error(),
		})
	}

	return c.JSON(application)
}