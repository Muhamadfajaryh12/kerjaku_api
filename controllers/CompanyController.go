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

func GetCompany(c *fiber.Ctx) error{
	var company []models.Company
	if err := databases.DB.Find(&company).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}
		return c.JSON(fiber.Map{"data":company})
}

func UpdateCompany(c *fiber.Ctx) error{
	var company models.Company
	id := c.Params("id")

	if err := databases.DB.Where("id = ?", id).First(&company); err == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Company tidak ditemukan"})
	}

	var companyUpdate models.Company
	if err := c.BodyParser(&companyUpdate) ; err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}

	photoUpload, err := c.FormFile("photo")
	if err == nil{
		photoPath,err := utils.UploadFile(photoUpload,"photo")
		if err!= nil{
			return c.Status(500).JSON(fiber.Map{"message":"Invalid "})
		}
		companyUpdate.Photo = photoPath
	}
	
	if  databases.DB.Model(&companyUpdate).Where("id = ?", id).Updates(&companyUpdate).RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message":"Gagal mengupdate company "})
	}

	return  c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil mengupdate company",
		"data": companyUpdate,
	})
}

func DeleteCompany(c *fiber.Ctx) error {
	var company models.Company
	id := c.Params("id")

	if err := databases.DB.Where("id = ?", id).First(&company); err == nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Company tidak ditemukan"})
	}

	databases.DB.Delete(&company,id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message:":"Berhasil menghapus company",
		"data":company,
	})
}

func DetailCompany(c *fiber.Ctx) error{
	id := c.Params("id")
	var company models.Company
	var vacancy []models.Vacancy

	if err := databases.DB.Where("id = ? ", id).First(&company).Error ; err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message":"Data tidak ditemukan"})
	}

	 databases.DB.Where("id_company", id).Find(&vacancy)

	 response := models.ICompanyVacancy{
		ID:          company.ID,
        CompanyName: company.CompanyName,
        CompanyType: company.CompanyType,
        Location:    company.Location,
        Size:        company.Size,
        Photo:       company.Photo,
        Description: company.Description,
		Vacancy: vacancy,
	 }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Detail company",
		"data":response,
	})
}

func SearchCompany(c *fiber.Ctx) error{
	query := c.Query("s")
	var company []models.Company

	if err := databases.DB.Where(`company_name LIKE ?`,"%"+query+"%").Find(&company); err != nil{
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":company,
	})
}
