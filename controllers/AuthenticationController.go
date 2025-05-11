package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func Register(c *fiber.Ctx) error {
var user models.User
var input models.User
	if err := c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}
	
	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if err := databases.DB.Where("username = ?", input.Username).First(&user).Error; 
	err == nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username sudah ada",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	input.Password = string(hash)

	databases.DB.Create(&input)
	return c.Status(fiber.StatusCreated).JSON(fiber	.Map{
		"message":"Berhasil membuat akun",
	})
}

func Login(c *fiber.Ctx) error{
	var user models.User
	var input  models.User
	if err := c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if err := databases.DB.Where("username = ?", input.Username).First(&user).Error; 
	err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Username dan Password salah"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password)); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"Username dan Password salah"})
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid token")
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Login berhasil",
		"token":token,
	})
}
