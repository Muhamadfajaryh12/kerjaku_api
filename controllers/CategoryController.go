package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func CategoryController(c *fiber.Ctx) error {
	var result models.CategoryData

	databases.DB.Model(&models.Vacancy{}).
		Distinct("location").
		Pluck("location", &result.Locations)

	databases.DB.Model(&models.Vacancy{}).
		Distinct("type").
		Pluck("type", &result.Types)

	databases.DB.Model(&models.Vacancy{}).
		Distinct("category").
		Pluck("category", &result.Categories)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

func CategoryCompanyController(c *fiber.Ctx) error {
	var result models.CategoryCompanyData

	databases.DB.Model(&models.Company{}).
		Distinct("location").
		Pluck("location", &result.Location)

	databases.DB.Model(&models.Company{}).
		Distinct("company_type").
		Pluck("company_type", &result.CompanyType)


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

