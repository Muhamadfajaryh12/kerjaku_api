package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InsertProfile(c *fiber.Ctx) error{
	var profile models.Profile
	if err := c.BodyParser(&profile) ; err != nil{
		return c.Status(400).JSON(fiber.Map{"message":"invalid request"})
	}
	databases.DB.Create(&profile)
	return c.JSON(profile)
}

func GetProfile (c *fiber.Ctx) error{
	var profile models.Profile
	id := c.Params("id")
	if err := databases.DB.Where("id_user = ?" , id).First(&profile).Error; err!= nil{
		switch err{
		case gorm.ErrRecordNotFound:
		return	c.Status(404).JSON(fiber.Map{"message":"Not Found"})
		default :
		return c.Status(400).JSON(fiber.Map{"Message":"Internal server Error"})
	}
}
return c.JSON(profile)
}