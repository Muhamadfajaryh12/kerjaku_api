package routes

import (
	"kerjaku/controllers"
	"kerjaku/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api")


	api.Post("/users",controllers.Register)
	api.Post("/login",controllers.Login)

	protectedRoute := api.Use(middlewares.AuthorizationMiddleware())

	apiProfile := protectedRoute.Group("/profile")
	apiProfile.Post("",controllers.InsertProfile)
	apiProfile.Get("/:id",controllers.GetProfile)

	apiVacancy := protectedRoute.Group("/vacancy")
	apiVacancy.Post("/",controllers.InsertVacancy)
	apiVacancy.Get("/",controllers.GetVacancy)
	apiVacancy.Get("/search",controllers.SearchVacancy)
	apiVacancy.Get("/:id",controllers.DetailVacancy)
	apiVacancy.Put("/:id",controllers.UpdateVacancy)
	apiVacancy.Delete("/:id",controllers.DeleteVacancy)
	
	apiCompany := protectedRoute.Group("/company")
	apiCompany.Post("/",controllers.InsertCompany)
	apiCompany.Get("/",controllers.GetCompany)
	apiCompany.Get("/:id",controllers.DetailCompany)
	apiCompany.Put("/:id",controllers.UpdateCompany)
	apiCompany.Delete("/:id",controllers.DeleteCompany)

}