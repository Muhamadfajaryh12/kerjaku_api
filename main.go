package main

import (
	model "kerjaku/databases"
	"kerjaku/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app  := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))
	model.ConnectionDatabase();
	routes.SetupRoutes(app);
	log.Fatal(app.Listen(":3000"))
}
