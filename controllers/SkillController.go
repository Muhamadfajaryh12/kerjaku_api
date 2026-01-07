package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func InsertSkill(c *fiber.Ctx) error {
	var input models.SkillForm
	userID := c.Locals("user_id").(float64)
	if err := c.BodyParser(&input);err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":err.Error()})
	}

	skill := models.Skill{
		Skill: input.Skill,
		UserID: uint(userID),
	}
	
	if err := databases.DB.Create(&skill).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":"Berhasil menambahkan keterampilan",
		"data":skill,
	})
}