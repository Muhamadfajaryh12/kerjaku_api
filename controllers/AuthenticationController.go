package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"

	"github.com/gofiber/fiber/v2"
)


func Register(c *fiber.Ctx) error {
var user models.User
	if err := c.BodyParser(&user); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	
	if err := utils.ValidateStruct(c,&user); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if err := databases.DB.Where("username = ?", user.Username).First(&user).Error; 
	err == nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username sudah ada",
		})
	}

	databases.DB.Create(&user)
	return c.Status(fiber.StatusCreated).JSON(fiber	.Map{
		"message":"Berhasil membuat akun",
	})
}

func Login(c *fiber.Ctx) error{
	var user models.User

	if err := c.BodyParser(&user); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	

	if err := utils.ValidateStruct(c,&user); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if err := databases.DB.Where("username = ?", user.Username).First(&user).Error; 
	err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Username dan Password salah")
	}

	return c.JSON(user)
}
