package controllers

func InsertLanguage(c *fiber.Ctx) error {
	var input models.LanguageForm
	userID := c.Locals("user_id").(float64)

	if err:= c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	language := models.Language{
		Language:input.Language,
		Level:input.Level,
		UserID:int64(userID)
	}

	if err := databases.DB.Create(&language); err != nil{
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.SendStatus(201).JSON{fiber.Map{
		"message":"berhasil menambahkan bahasa",
		"data":language
	}}
}


func DeleteLanguage(c *fiber.Ctx) error {
	var language models.Language
	id := c.Params("id")

	if err:=databases.DB.Where("id = ?",id).First(&language); err != nil {
		return c.Status(fiber.StatusInternalServerError)
	}
	
	if err:= databases.DB.Delete(&language, id); err != nil {
		return c.Status(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":"berhasil menghapus bahasa",
		"id":id
	})
}