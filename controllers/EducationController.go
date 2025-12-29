package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func InsertEducation(c *fiber.Ctx) error {
	var input models.EducationForm

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err:= databases.DB.Create(&input); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(input)
}