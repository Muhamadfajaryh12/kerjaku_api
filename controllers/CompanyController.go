package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"

	"github.com/gofiber/fiber/v2"
)

func InsertCompany(c *fiber.Ctx) error {
	var company models.Company
	if err := c.BodyParser(&company) ; err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}
	photoUpload, err := c.FormFile("photo")
	if err == nil{
		photoPath,err := utils.UploadFile(photoUpload,"photo")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		company.Photo = photoPath
	}
	databases.DB.Create(&company)
	return c.JSON(company)
}

func GetCompany (c *fiber.Ctx) error{
	var company []models.Company
	if err := databases.DB.Find(&company).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}
		return c.JSON(fiber.Map{"data":company})

}