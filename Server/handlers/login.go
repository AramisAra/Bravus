package handlers

import (
	"time"

	"github.com/AramisAra/GroomingApp/config"
	"github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	login := models.Login{}
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var client models.Client
	result := database.Database.Db.Find(&client, "Email = ?", login.Email)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(login.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid Password"})
	}

	claims := jwt.MapClaims{
		"ID":    client.ID,
		"email": client.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiry set to 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
	}

	// Return the token
	return c.JSON(models.LoginResponse{Token: t})

}

func Protected(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token) // Ensure 'user' matches the ContextKey in middleware
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)
	return c.SendString("Welcome 👋 " + email + " to the Bravus dashboard")
}
