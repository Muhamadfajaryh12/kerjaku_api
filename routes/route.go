package routes

import (
	"kerjaku/controllers"
	"kerjaku/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	app.Static("uploads","./uploads")

	api := app.Group("/api")


	api.Post("/register",controllers.Register)
	api.Post("/login",controllers.Login)


	api.Get("vacancy",controllers.GetVacancy)
	api.Get("vacancy/company/:id",controllers.GetVacancyCompany)
	api.Get("vacancy/:id",controllers.DetailVacancy)

	api.Get("category",controllers.CategoryController)
	api.Get("category/company",controllers.CategoryCompanyController)

	api.Get("company",controllers.GetCompany)
	api.Get("company/search",controllers.SearchCompany)
	api.Get("company/:id",controllers.DetailCompany)
	api.Get("dashboard",controllers.DashboardApplication)

	
	protectedRoute := api.Use(middlewares.AuthorizationMiddleware())
	
	apiExperience := protectedRoute.Group("/experience")
	apiExperience.Get("/:id",controllers.DetailExperience)
	apiExperience.Post("/", controllers.InsertExperience)
	apiExperience.Put("/:id",controllers.UpdateExperience)
	apiExperience.Delete("/:id",controllers.DeleteExperience)

	apiProfile := protectedRoute.Group("/profile")
	apiProfile.Get("/",controllers.GetProfile)
	apiProfile.Post("/",controllers.InsertProfile)
	apiProfile.Put("/:id", controllers.UpdateProfile)

	apiVacancy := protectedRoute.Group("/vacancy")
	apiVacancy.Post("/",controllers.InsertVacancy)
	apiVacancy.Put("/:id",controllers.UpdateVacancy)
	apiVacancy.Delete("/:id",controllers.DeleteVacancy)

	
	apiCompany := protectedRoute.Group("/company")
	apiCompany.Post("/",controllers.InsertCompany)
	apiCompany.Put("/:id",controllers.UpdateCompany)
	apiCompany.Delete("/:id",controllers.DeleteCompany)

	apiApplication := protectedRoute.Group("/application")
	apiApplication.Post("/",controllers.InsertApplication)
	apiApplication.Get("/",controllers.GetApplication)
	apiApplication.Get("/:id",controllers.GetDetailApplication)
	apiApplication.Put("/:id",controllers.UpdateApplication)
	apiApplication.Delete("/:id",controllers.DeleteApplication)


	apiEducation := protectedRoute.Group("/education")
	apiEducation.Post("/", controllers.InsertEducation)
	apiEducation.Delete("/:id",controllers.DeleteEducation)

	apiLanguage := protectedRoute.Group("/language")
	apiLanguage.Post("/",controllers.InsertLanguage)
	apiLanguage.Delete("/:id",controllers.DeleteLanguage)

	apiCertification := protectedRoute.Group("/certification")
	apiCertification.Post("/",controllers.InsertCertification)
	apiCertification.Delete("/:id",controllers.DeleteCertification)

	apiSkill := protectedRoute.Group("/skill")
	apiSkill.Post("/",controllers.InsertSkill)
	apiSkill.Delete("/:id",controllers.DeleteSkill)
}