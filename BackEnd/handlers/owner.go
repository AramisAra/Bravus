package handlers

import (
	"time"

	"github.com/AramisAra/BravusBackend/config"
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Register Owner func
func RegisterOwner(c *fiber.Ctx) error {
	registration := new(dbmodels.RegisterRequestOwner)
	if err := c.BodyParser(registration); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the email already exists
	var existingOwner dbmodels.Owner
	if err := database.Database.Db.Where("email = ?", registration.Email).First(&existingOwner).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already registered"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not hash password"})
	}

	// Create the user
	owner := dbmodels.Owner{
		Full_Name: registration.Name,
		Email:     registration.Email,
		Phone:     registration.Phone,
		Password:  string(hashedPassword),
		Career:    registration.Career,
	}

	// Save the user in the database
	result := database.Database.Db.Create(&owner)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(result.Error)
	}

	return c.Status(fiber.StatusCreated).JSON(owner)
}

// Login Owner Func
func LoginOwner(c *fiber.Ctx) error {
	login := dbmodels.Login{}
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Checkes if the email is in the database
	var owner dbmodels.Owner
	result := database.Database.Db.Find(&owner, "Email = ?", login.Email)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	// Checkes hash password
	if err := bcrypt.CompareHashAndPassword([]byte(owner.Password), []byte(login.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid Password"})
	}

	// Creates the Claims
	claims := jwt.MapClaims{
		"ID":    owner.ID,
		"email": owner.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expiry set to 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
	}

	// Return the token
	return c.JSON(dbmodels.LoginResponse{Token: t})
}

// List all owners
func ListOwners(c *fiber.Ctx) error {
	owners := []dbmodels.Owner{}

	database.Database.Db.Find(&owners)
	responseOwner := []dbmodels.OwnerSerializer{}

	for _, owner := range owners {
		responseOwner = append(responseOwner, dbmodels.CreateOwnerResponse(owner))
	}

	return c.Status(200).JSON(responseOwner)
}
