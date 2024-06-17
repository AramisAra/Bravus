package handlers

import (
	database "github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateOwner(c *fiber.Ctx) error {
	var owner models.Owner

	if err := c.BodyParser(&owner); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&owner)
	responseOwner := database.CreateOwnerResponse(owner)

	return c.Status(200).JSON(responseOwner)
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
