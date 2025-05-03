package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InsertProfile(c *fiber.Ctx) error{
	var profile models.Profile
	if err := c.BodyParser(&profile) ; err != nil{
		return c.Status(400).JSON(fiber.Map{"message":"invalid request"})
	}

	cvUpload, err := c.FormFile("cv")
	
	if err == nil{
		cvPath,err := utils.UploadFile(cvUpload,"cv")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		profile.CV = cvPath
	}
	
	photoUpload, err := c.FormFile("photo")
	if err == nil{
		photoPath,err := utils.UploadFile(photoUpload,"photo")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		profile.Photo = photoPath
	}

	databases.DB.Create(&profile)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil membuat profile",
		"data":profile,
	})

}

func GetProfile(c *fiber.Ctx) error{
	var profile models.Profile
	id := c.Params("id")
	if err := databases.DB.Where("id_user = ?" , id).First(&profile).Error; err!= nil{
		switch err{
		case gorm.ErrRecordNotFound:
		return	c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Not Found"})
		default :
		return c.SendStatus(fiber.ErrBadRequest.Code)
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data":profile})
}

func UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var profile models.Profile

	if err := c.BodyParser(&profile) ; err != nil{
		return c.Status(400).JSON(fiber.Map{"message":"invalid request"})
	}
	
	cvUpload, err := c.FormFile("cv")
	
	if err == nil && cvUpload != nil{
		cvPath,err := utils.UploadFile(cvUpload,"cv")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		profile.CV = cvPath
	}
	
	photoUpload, err := c.FormFile("photo")

	if err == nil && photoUpload != nil{
		photoPath,err := utils.UploadFile(photoUpload,"photo")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		profile.Photo = photoPath
	}


	if databases.DB.Model(&profile).Where("id = ?", id).Updates(&profile).RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Gagal mengupdate profile "})
	}

	return  c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil mengupdate profile",
		"data": profile,
	})
}