package main

import (
	model "kerjaku/databases"
	"kerjaku/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app  := fiber.New()
	model.ConnectionDatabase();
	routes.SetupRoutes(app);
	log.Fatal(app.Listen(":3000"))
}
