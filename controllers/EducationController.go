package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func InsertEducation(c *fiber.Ctx) error {
	var input models.EducationForm
	userID := c.Locals("user_id").(float64)
	if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	
	education := models.Education{
		EducationName: input.EducationName,
		Major:         input.Major,
		GraduateDate:  input.GraduateDate,
		Level:input.Level,
		UserID:       uint(userID),
	}

	if err:= databases.DB.Create(&education).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":err.Error(),
		})	}	


	return c.Status(201).JSON(fiber.Map{
		"message":"berhasil menambahkan pendidikan",
		"data":education,
	})
}


func DeleteEducation(c *fiber.Ctx) error {
	var education models.Education
	id := c.Params("id")

	if err := databases.DB.Where("id = ?",id).First(&education).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":err.Error(),
		})
	}

	if err := databases.DB.Delete(&education).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":err.Error(),
		})	
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"berhasil menghabpus pendidikan",
		"id":id,
	})
}