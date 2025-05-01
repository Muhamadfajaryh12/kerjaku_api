package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"

	"github.com/gofiber/fiber/v2"
)

func InsertVacancy(c *fiber.Ctx) error {
	var vacancy models.Vacancy
	if err := c.BodyParser(&vacancy) ; err != nil {
		return c.Status(400).JSON(fiber.Map{"message":err.Error()})
	}

	databases.DB.Create(&vacancy)

	databases.DB.Preload("Company").First(&vacancy, vacancy.ID)
	return c.JSON(vacancy)
}

func SearchVacancy(c *fiber.Ctx) error{
	query := c.Query("s")
	var vacancy []models.Vacancy
	
	if err := databases.DB.Where("name_vacancy LIKE ?", "%"+query+"%").Find(&vacancy).Error ; err != nil{
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":vacancy,
	})

}

func GetVacancy(c *fiber.Ctx) error {
	var vacancy []models.Vacancy
	if err := databases.DB.Find(&vacancy).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}
	return c.JSON(fiber.Map{"data":vacancy})

}

func DetailVacancy(c *fiber.Ctx) error{
	id := c.Params("id")
	var vacancy models.Vacancy

	if err := databases.DB.Where("id = ? ", id).First(&vacancy).Error; err != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message":"Vacancy tidak ditemukan",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Detail Vacancy",
		"data":vacancy,
	})
}


func UpdateVacancy(c *fiber.Ctx) error{
	id := c.Params("id")
	var vacancy models.Vacancy
	var input models.Vacancy

	if err := c.BodyParser(&input); err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := utils.ValidateStruct(c, &input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if databases.DB.Model(&vacancy).Where("id = ?",id).Updates(&input).RowsAffected == 0{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Gagal mengubah vacancy"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil mengubah vacancy",
		"data":vacancy,
	})
}

func DeleteVacancy(c* fiber.Ctx) error{
	id := c.Params("id")
	var vacancy models.Vacancy

	if err := databases.DB.Where("id = ?", id).First(&vacancy); err == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Vacancy tidak ditemukan"})
	}

	databases.DB.Delete(&vacancy,id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message:":"Berhasil menghapus company",
		"data":vacancy,
	})

}