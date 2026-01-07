package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func InsertCertification(c *fiber.Ctx) error{
	var input models.CertificationForm
	userID := c.Locals("user_id").(float64)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":err.Error()});
	}

	certification := models.Certification{
		CertificationName: input.CertificationName,
		Publisher: input.Publisher,
		EffectiveDate: input.EffectiveDate,
		UserID: uint(userID),

	} 

	if err := databases.DB.Create(&certification).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()});
	}
	
	return c.Status(201).JSON(fiber.Map{
		"message":"Berhasil menambahkan sertifikat",
		"data":certification,
	})
}