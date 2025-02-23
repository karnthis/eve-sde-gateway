package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"sde/database"
)

func main() {
	database.InitDB()
	app := fiber.New()
	app.Use(cors.New())
	setupRoutes(app)

	app.Listen(":3000")
}
