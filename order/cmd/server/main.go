package main

import (
	"fast_food_order/internals/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := createConfig()

	app := fiber.New(*config)

	app.Get("/healthy/:Key", func(c *fiber.Ctx) error {
		key := c.Params("Key")

		if key != "" {
			return fiber.NewError(fiber.StatusBadRequest, "Error Key invalid")
		}

		return c.JSON(fiber.Map{
			"running": true,
		})
	})

	app.Listen(":3030")
}

func createConfig() *fiber.Config {
	return &fiber.Config{
		Prefork:      true,
		ErrorHandler: middleware.GlobalErrorHandler,
		AppName:      "Fast Food GO",
	}
}
