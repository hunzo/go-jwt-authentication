package main

import (
	"api/database"
	"api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)

	app.Listen(":8080")

}
