package handlers

import (
	"github.com/gofiber/fiber/v2"
	"sde/database"
	"strconv"
)

func Traits(c *fiber.Ctx) error {
	submission := c.Query("selected")
	value, err := strconv.Atoi(submission)
	if err != nil {
		// ... handle error
		panic(err)
	}

	data := database.ReadTraits(value)
	return c.Status(200).JSON(data)
}
