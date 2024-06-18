package handlers

import (
	database "github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterOwner(c *fiber.Ctx) error {
	registration := new(models.RegisterRequestOwner)
	if err := c.BodyParser(registration); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Check if the email already exists
	var existingOwner models.Client
	if err := database.Database.Db.Where("email = ?", registration.Email).First(&existingOwner).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already registered"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not hash password"})
	}

	// Create the user
	owner := models.Owner{
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
	owners := []models.Owner{}

	database.Database.Db.Find(&owners)
	responseOwner := []database.OwnerSerializer{}

	for _, owner := range owners {
		responseOwner = append(responseOwner, database.CreateOwnerResponse(owner))
	}

	return c.Status(200).JSON(responseOwner)
}

func GetOwner(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	owner := models.Owner{}

	if err := database.FindOwner(parsedId, &owner); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseOwner := database.CreateOwnerResponse(owner)

	return c.Status(200).JSON(responseOwner)
}

func UpdateOwner(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	owner := models.Owner{}
	if err := database.FindOwner(parsedId, &owner); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateOwner database.UpdateOwnerInput

	if err := c.BodyParser(&updateOwner); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	owner.Full_Name = updateOwner.Full_Name
	owner.Phone = updateOwner.Phone
	owner.Email = updateOwner.Email
	owner.Career = updateOwner.Career

	responseOwner := database.CreateOwnerResponse(owner)
	return c.Status(200).JSON(responseOwner)
}

func DeleteOwner(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	owner := models.Owner{}
	if err := database.FindOwner(parsedId, &owner); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&owner).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Owner Was Deleted")
}

func GetAppointmentOwner(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var owner models.Owner

	database.Database.Db.Preload("Appointments").Find(&owner)

	response := database.CreateOwnerResponse(owner)

	return c.Status(200).JSON(response)
}
