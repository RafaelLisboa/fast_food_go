package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GlobalErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var message string

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	} else {
		message = err.Error()
	}

	log.Error("Error: %s", err)

	return c.Status(code).JSON(fiber.Map{
		"error":   true,
		"message": message,
	})
}
