package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// This is for testing only
func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token) // Ensure 'user' matches the ContextKey in middleware
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)
	return c.SendString("Welcome 👋 " + email + " to the Bravus dashboard")
}