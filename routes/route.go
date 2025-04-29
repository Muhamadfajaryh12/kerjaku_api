package routes

import (
	"kerjaku/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	apiCompany := api.Group("/company")
	api.Post("/users",controllers.Register)
	api.Post("/login",controllers.Login)
	api.Post("/profile",controllers.InsertProfile)
	api.Get("/profile/:id",controllers.GetProfile)
	api.Post("/vacancy",controllers.InsertVacancy)
	api.Get("/vacancy",controllers.GetVacancy)
	apiCompany.Post("/",controllers.InsertCompany)
	apiCompany.Get("/",controllers.GetCompany)
	apiCompany.Put("/:id",controllers.UpdateCompany)
	apiCompany.Delete("/:id",controllers.DeleteCompany)

}