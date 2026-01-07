package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func InsertLanguage(c *fiber.Ctx) error {
	var input models.LanguageForm
	userID := c.Locals("user_id").(float64)

	if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":err.Error(),
		})
	}

	language := models.Language{
		Language: input.Language,
		Level:    input.Level,
		UserID:   uint(userID),
	}

	if err := databases.DB.Create(&language).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
		
	return c.Status(201).JSON(fiber.Map{
		"message": "berhasil menambahkan bahasa",
		"data":    language,
	})
}

func DeleteLanguage(c *fiber.Ctx) error {
	var language models.Language
	id := c.Params("id")

	if err := databases.DB.Where("id = ?", id).First(&language); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := databases.DB.Delete(&language, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "berhasil menghapus bahasa",
		"id":      id,
	})
}