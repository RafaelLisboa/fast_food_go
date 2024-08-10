package server

import (
	"fast_food_order/internals/handler"
	"fast_food_order/internals/middleware"

	"github.com/gofiber/fiber/v2"
)

func StartServer() {

	config := createConfig()
	handler := handler.NewOrderHandler()

	app := fiber.New(*config)

	app.Get("/healthy/", healthyHandler)

	app.Get("/message", handler.SendMessage)

	app.Listen(":3030")
}

func createConfig() *fiber.Config {
	return &fiber.Config{
		Prefork:      false,
		ErrorHandler: middleware.GlobalErrorHandler,
		AppName:      "Fast Food GO",
	}
}

func healthyHandler(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{
		"running": true,
	})

}
