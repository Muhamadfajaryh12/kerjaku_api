package controllers

import (
	"errors"
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func Register(c *fiber.Ctx) error {
var user models.User
var input models.User
	if err := c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":err.Error(),
		})
	}
	
	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":err,
		})
	}

	if err := databases.DB.Where("email = ?", input.Email).First(&user).Error; 
	err == nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email sudah ada",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
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
	var input  models.LoginForm

	var company models.Company

	if err := c.BodyParser(&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":err.Error(),
		})	}

	if err := utils.ValidateStruct(c,&input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if err := databases.DB.Where("email = ?", input.Email).First(&user).Error; 
	err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"email dan Password salah"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(input.Password)); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message":"email dan Password salah"})
	}

	isComplete := true

	if(user.Role == "pencaker"){
		var profile models.Profile
		if err := databases.DB.Where("user_id",user.ID).First(&profile).Error; err != nil{
			if errors.Is(err, gorm.ErrRecordNotFound){
				isComplete = false
			}else{
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message":err.Error(),
				})
			}
		}

		var certification []models.Certification
		databases.DB.Where("user_id",user.ID).First(&certification)

		if len(certification) == 0 {
			isComplete = false
		}

		var experience []models.Experience
		databases.DB.Where("user_id",user.ID).First(&experience)

		if len(experience) == 0 {
			isComplete = false
		}

		var language []models.Language
		databases.DB.Where("user_id",user.ID).First(&language)

		if len(language) == 0 {
			isComplete = false
		}

		var education []models.Education
	 	databases.DB.Where("user_id",user.ID).First(&education)

		if len(education) == 0 {
			isComplete = false
		}

		var skill []models.Skill
		 databases.DB.Where("user_id",user.ID).First(&skill)

		if len(skill) == 0 {
			isComplete = false
		}
	}else{
		
		if err := databases.DB.Where("user_id",user.ID).First(&company).Error; err != nil{
			if errors.Is(err, gorm.ErrRecordNotFound){
				isComplete = false
			}else{
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message":err.Error(),
				})
			}
		}
	}

	token, err := utils.GenerateToken(user.ID,isComplete)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).SendString("Invalid token")
	}


	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Login berhasil",
		"token":token,
		"id":user.ID,
		"is_complete":isComplete,
	})
}
