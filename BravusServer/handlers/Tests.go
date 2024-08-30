package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// This is for testing only
func Protected(c *fiber.Ctx) error {
	return c.SendString("Welcome 👋 to the Bravus dashboard")
}
