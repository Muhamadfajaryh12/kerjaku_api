package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func InsertApplication(c *fiber.Ctx) error {
	var application models.Application
	var input models.Application

	if err := c.BodyParser(&input); err != nil{
		return c.Status(400).JSON(fiber.Map{"message":err.Error()})
	}
	input.Status = "Menunggu"
	databases.DB.Create(&input).Find(&application)
	
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Berhasil melamar pekerjaan",
		"data":application,
	})
}

func GetApplication(c *fiber.Ctx) error {
	userId := c.Query("user")
	vacancyId := c.Query("vacancy")
	var application []models.Application

	if userId != "" {
		databases.DB.Where("id_user", userId).Find(&application)
	}
	
	
	if vacancyId != "" {
		databases.DB.Where("id_vacancy",vacancyId).Find(&application)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"List application",
		"data":application,
	})
}

func UpdateApplication(c *fiber.Ctx) error{
	id:= c.Params("id")
	var input models.Application
	var application models.Application
	
	if err := c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":"Bad Request",
		})
	}

	if  databases.DB.Model(&application).Where("id = ?", id).Updates(&input).RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Gagal mengupdate vacancy "})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message":"Berhasil mengupdate vacancy","data":application})
}

func DeleteApplication(c *fiber.Ctx) error {
	id := c.Params("id")	
	var application models.Application

	if err := databases.DB.Delete(&application,id); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":"Internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil menghapus application",
		"data":application,
	})
}