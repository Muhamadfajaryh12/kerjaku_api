package routes

import (
	"kerjaku/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/users",controllers.Register)
	api.Post("/login",controllers.Login)
	api.Post("/profile",controllers.InsertProfile)
	api.Get("/profile/:id",controllers.GetProfile)
	api.Post("/company",controllers.InsertCompany)
	api.Get("/company",controllers.GetCompany)
	api.Post("/vacancy",controllers.InsertVacancy)
	api.Get("/vacancy",controllers.GetVacancy)

}