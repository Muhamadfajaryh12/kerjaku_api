package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func InsertVacancy(c *fiber.Ctx) error {
	var vacancy models.Vacany
	if err := c.BodyParser(&vacancy) ; err != nil {
		return c.Status(400).JSON(fiber.Map{"message":err.Error()})
	}

	databases.DB.Create(&vacancy)

	databases.DB.Preload("Company").First(&vacancy, vacancy.ID)
	return c.JSON(vacancy)
}


func GetVacancy (c *fiber.Ctx) error {
	var vacancy []models.Vacany
	if err := databases.DB.Preload("Company").Find(&vacancy).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}
	return c.JSON(fiber.Map{"data":vacancy})

}