package controllers

import (
	"kerjaku/databases"
	"kerjaku/models"
	"kerjaku/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func InsertVacancy(c *fiber.Ctx) error {
	var vacancy models.Vacancy
	if err := c.BodyParser(&vacancy) ; err != nil {
		return c.Status(400).JSON(fiber.Map{"message":err.Error()})
	}

	databases.DB.Create(&vacancy)

	databases.DB.Preload("Company").First(&vacancy, vacancy.ID)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Berhasil membuat lowongan",
		"data":vacancy,
	})
}


func SearchFilterVacancy(c *fiber.Ctx) error{
	var filter models.VacancyFilter
	c.QueryParser(&filter)
	
	var vacancy []models.Vacancy
	query := databases.DB.
        Joins("JOIN companies ON companies.id = vacancies.id_company").
        Preload("Company")

    if filter.Category != "" {
        categories := strings.Split(filter.Category, ",")
        if len(categories) > 1 {
            query = query.Where("vacancies.category IN ?", categories)
        } else {
            query = query.Where("vacancies.category = ?", filter.Category)
        }
    }

    if filter.Location != "" {
        locations := strings.Split(filter.Location, ",")
        if len(locations) > 1 {
            query = query.Where("vacancies.location IN ?", locations)
        } else {
            query = query.Where("vacancies.location = ?", filter.Location)
        }
    }

    if filter.Type != "" {
        types := strings.Split(filter.Type, ",")
        if len(types) > 1 {
            query = query.Where("vacancies.type IN ?", types)
        } else {
            query = query.Where("vacancies.type = ?", filter.Type)
        }
    }

    if filter.Status != "" {
        query = query.Where("vacancies.status = ?", filter.Status)
    }

	if filter.Search != ""{
		query = query.Where("vacancies.name_vacancy LIKE ?", "%"+filter.Search+"%")
	} 

    if err := query.Find(&vacancy).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to fetch vacancies",
        })
    }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":vacancy,
	})

}

func GetVacancy(c *fiber.Ctx) error {
	var vacancy []models.Vacancy
	// var response []models.IVacancy
	if err := databases.DB.Preload("Company").Find(&vacancy).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}
	// for _, vacancy := range vacancy {
    //     var company models.Company
    //     if err := databases.DB.Where("id = ?", vacancy.IDCompany).First(&company).Error; err != nil {
    //         company = models.Company{}
    //     }
		
    //     response = append(response, models.IVacancy{
    //         ID:          vacancy.ID,
    //         NameVacancy: vacancy.NameVacancy,
    //         Description: vacancy.Description,
    //         Location:    vacancy.Location,
    //         Qty:         vacancy.Qty,
    //         Salary:      vacancy.Salary,
    //         DateEnd:     vacancy.DateEnd,
    //         DateStart:   vacancy.DateStart,
    //         Status:      vacancy.Status,
    //         IDCompany:   vacancy.IDCompany,
    //         Company:     company,
    //     })
    // }
	return c.JSON(fiber.Map{"data":vacancy})

}

func DetailVacancy(c *fiber.Ctx) error{
	id := c.Params("id")
	var vacancy models.Vacancy
		if err := databases.DB.Preload("Company").Where("vacancies.id = ?",id).Find(&vacancy).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal mengambil data"})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Detail Vacancy",
		"data":vacancy,
	})
}


func UpdateVacancy(c *fiber.Ctx) error{
	id := c.Params("id")
	var vacancy models.Vacancy
	var input models.Vacancy

	if err := c.BodyParser(&input); err != nil{
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := utils.ValidateStruct(c, &input); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
		})
	}

	if databases.DB.Model(&vacancy).Where("id = ?",id).Updates(&input).RowsAffected == 0{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message":"Gagal mengubah vacancy"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"Berhasil mengubah vacancy",
		"data":vacancy,
	})
}

func DeleteVacancy(c* fiber.Ctx) error{
	id := c.Params("id")
	var vacancy models.Vacancy

	if err := databases.DB.Where("id = ?", id).First(&vacancy); err == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message":"Vacancy tidak ditemukan"})
	}

	databases.DB.Delete(&vacancy,id)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message:":"Berhasil menghapus company",
		"data":vacancy,
	})

}