package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

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

func ListOwners(c *fiber.Ctx) error {
	owners := []dbmodels.Owner{}

	database.Database.Db.Find(&owners)
	responseOwner := []dbmodels.OwnerSerializer{}

	for _, owner := range owners {
		responseOwner = append(responseOwner, dbmodels.CreateOwnerResponse(owner))
	}

	return c.Status(200).JSON(responseOwner)
}
