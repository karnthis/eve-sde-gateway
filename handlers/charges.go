package handlers

import (
	"github.com/gofiber/fiber/v2"
	"sde/database"
)

func Modules(c *fiber.Ctx) error {
	data := database.ReadModules()
	return c.Status(200).JSON(data)
}
