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

func GetEducation(c *fiber.Ctx) error {
	var education []models.EducationForm
	
	if err := databases.DB.Find(&education); err !nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(education)
}

func DeleteEducation(c *fiber.Ctx) error {
	var education models.Education
	id := c.Params("id")

	if err := databases.DB.Where("id = ?",id).First(&education); err != nil {
		return c.Status(fiber.StatusInternalServerError)
	}

	if err := databases.DB.Delete(&education, id); err != nil {
		return c.Status(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"berhasil menghabpus pendidikan",
		"id":id
	})
}