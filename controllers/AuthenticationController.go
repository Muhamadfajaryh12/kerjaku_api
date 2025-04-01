package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
var user models.User
	if err := c.BodyParser(&user); err != nil{
		return c.Status(400).JSON(fiber.Map{"Message":"Invalid request"})
	}
	databases.DB.Create(&user)
	return c.JSON(user)
}
