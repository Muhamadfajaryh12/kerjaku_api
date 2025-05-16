package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"

	"github.com/gofiber/fiber/v2"
)

func InsertExperience(c *fiber.Ctx) error {
	var experience models.Experience
	if err := c.BodyParser(&experience) ; err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":err,
		})
	}

	if err := utils.ValidateStruct(c,&experience); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	databases.DB.Create(&experience)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Berhasil menambahkan experience",
	})
}

func DetailExperience(c *fiber.Ctx) error{
	id:= c.Params("id")
	var experience models.Experience
	if err := databases.DB.Where("id = ?", id).First(&experience) ; err == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Experience tidak ditemukan"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":experience,
	})
}

func UpdateExperience(c *fiber.Ctx) error {
	var input models.Experience
	var experience models.Experience
	id:= c.Params("id")

	if err := databases.DB.Where("id = ?", id).First(&experience); err == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Experience tidak ditemukan"})
	}


	if err := c.BodyParser(&input) ; err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if  databases.DB.Model(&experience).Updates(&input).RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Gagal mengupdate experience "})
	}

	return  c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil mengupdate experience",
		"data": input,
	})
}

func DeleteExperience(c *fiber.Ctx) error {
	var experience models.Experience
	id := c.Params("id")

	if err := databases.DB.Where("id = ?", id).First(&experience); err == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"experience tidak ditemukan"})
	}

	databases.DB.Delete(&experience,id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message:":"Berhasil menghapus experience",
		"id":id,
	})
}
