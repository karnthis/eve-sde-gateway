package main

import (
	"github.com/gofiber/fiber/v2"
	"sde/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/ships", func(c *fiber.Ctx) error {
		return handlers.Ships(c)
	})

	app.Get("/modules", func(c *fiber.Ctx) error {
		return handlers.Modules(c)
	})

	app.Get("/charges", func(c *fiber.Ctx) error {
		return handlers.Charges(c)
	})

	app.Get("/ship_traits", func(c *fiber.Ctx) error {
		return handlers.Traits(c)
	})

	app.Get("/module_traits", func(c *fiber.Ctx) error {
		return handlers.Traits(c)
	})

	app.Get("/charge_traits", func(c *fiber.Ctx) error {
		return handlers.Traits(c)
	})

}
