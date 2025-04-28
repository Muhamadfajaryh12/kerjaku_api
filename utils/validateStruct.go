// utils/validation.go
package utils

import (
	"kerjaku/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidateStruct(c *fiber.Ctx, structData interface{}) []*models.IError {
	validate := validator.New()

	if err := validate.Struct(structData); err != nil {
		var errors []*models.IError
		for _, err := range err.(validator.ValidationErrors) {
			var el models.IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}

		return errors
	}

	return nil
}
