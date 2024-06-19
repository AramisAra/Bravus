package handlers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

// Login Test Route
func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token) // Ensure 'user' matches the ContextKey in middleware
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	return c.SendString("Welcome 👋 " + email + " to the Bravus dashboard")
}
