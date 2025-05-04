package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InsertProfile(c *fiber.Ctx) error{
	var input models.Profile
	
	if err := c.BodyParser(&input) ; err != nil{
		return c.Status(400).JSON(fiber.Map{"message":"invalid request"})
	}

	cvUpload, err := c.FormFile("cv")
	
	if err == nil{
		cvPath,err := utils.UploadFile(cvUpload,"cv")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		input.CV = cvPath
	}
	
	photoUpload, err := c.FormFile("photo")
	if err == nil{
		photoPath,err := utils.UploadFile(photoUpload,"photo")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		input.Photo = photoPath
	}

	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}


	databases.DB.Create(&input)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil membuat profile",
		"data":input,
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
	var input models.Profile

	if err := c.BodyParser(&input) ; err != nil{
		return c.Status(400).JSON(fiber.Map{"message":"invalid request"})
	}
	
	databases.DB.Where("id = ?", id).First(&profile)

	cvUpload, err := c.FormFile("cv")
	
	if err == nil && cvUpload != nil{
		utils.DeleteFile(strings.ReplaceAll(profile.CV, "/", "\\"))

		cvPath,err := utils.UploadFile(cvUpload,"cv")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		input.CV = cvPath
	}
	
	photoUpload, err := c.FormFile("photo")

	if err == nil && photoUpload != nil{
		utils.DeleteFile(strings.ReplaceAll(profile.Photo, "/", "\\"))
	

		photoPath,err := utils.UploadFile(photoUpload,"photo")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		input.Photo = photoPath
	}

	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if databases.DB.Model(&profile).Updates(&input).RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Gagal mengupdate profile "})
	}

	return  c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil mengupdate profile",
		"data": input,
	})
}