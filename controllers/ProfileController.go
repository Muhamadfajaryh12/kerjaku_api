package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func InsertProfile(c *fiber.Ctx) error{
	var input models.Profile
	userID := c.Locals("user_id").(float64)
	if err := c.BodyParser(&input) ; err != nil{
		return c.Status(400).JSON(fiber.Map{"message":"invalid request"})
	}

	profile := models.Profile{
		Name: input.Name,
		Address: input.Address,
		Birth: input.Birth,
		Phone: input.Phone,
		Email: input.Email,
		Gender: input.Gender,
		Summary: input.Summary,
		UserID: uint(userID),
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

	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err,
		})
	}

	if err := databases.DB.Create(&profile).Error; err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Berhasil membuat profile",
		"data":profile,
	})

}

func GetProfile(c *fiber.Ctx) error{
	userID := c.Locals("user_id").(float64)

	var education []models.Education
	if err := databases.DB.Where("user_id = ?", userID).Find(&education).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	var language []models.Language
	if err := databases.DB.Where("user_id = ?",userID).Find(&language).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	var certification []models.Certification
	if err := databases.DB.Where("user_id = ?",userID).Find(&certification).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	var skill []models.Skill
	if err := databases.DB.Where("user_id = ?",userID).Find(&skill).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	var experience []models.Experience
	if err := databases.DB.Where("user_id = ?",userID).Find(&experience).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	var profiles models.Profile
	if err := databases.DB.Where("user_id = ?",userID).Find(&profiles).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":err.Error()})
	}

	profile := models.ProfileResponse{
		Education: education,
		Language: language,
		Certification: certification,
		Skill: skill,
		Experience: experience,
		Profile: profiles,
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
	}else{
		input.CV = profile.CV
	}
	
	photoUpload, err := c.FormFile("photo")

	if err == nil && photoUpload != nil{
		utils.DeleteFile(strings.ReplaceAll(profile.Photo, "/", "\\"))
	

		photoPath,err := utils.UploadFile(photoUpload,"photo")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		input.Photo = photoPath
	} else {
		input.Photo = profile.Photo
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