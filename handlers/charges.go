package handlers

import (
	"github.com/gofiber/fiber/v2"
	"sde/database"
)

func Charges(c *fiber.Ctx) error {
	data := database.ReadCharges()
	return c.Status(200).JSON(data)
}
