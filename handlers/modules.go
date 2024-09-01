package handlers

import (
	"github.com/gofiber/fiber/v2"
	"sde/database"
)

func Ships(c *fiber.Ctx) error {
	data := database.ReadShips()
	return c.Status(200).JSON(data)
}
