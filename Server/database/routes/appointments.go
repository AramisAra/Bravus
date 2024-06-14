package routes

import (
	"time"

	"github.com/AramisAra/GroomingApp/database"
	"github.com/AramisAra/GroomingApp/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAppointment(c *fiber.Ctx) error {
	userid := c.Params("uuidUser")
	if userid == "" {
		userid = c.Query("uuidUser")
	}
	ownerid := c.Params("uuidOwner")
	if ownerid == "" {
		ownerid = c.Query("uuidOwner")
	}
	animalid := c.Params("uuidAnimal")
	if animalid == "" {
		animalid = c.Query("uuidAnimal")
	}

	if !isValidUUID(userid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	if !isValidUUID(ownerid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	if !isValidUUID(animalid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var client models.Client
	var owner models.Owner
	var animal models.Animals

	database.Database.Db.Find(&client, "id = ?", userid)
	database.Database.Db.Find(&owner, "id = ?", ownerid)
	database.Database.Db.Find(&animal, "id = ?", animalid)

	clientID := client.ID
	ownerID := owner.ID
	animalID := animal.ID

	var appointment models.Appointment

	appointment.ClientID = clientID
	appointment.OwnerID = ownerID
	appointment.AnimalID = animalID

	layoutUS := "01/01/2024"
	var err error
	appointment.Date, err = time.Parse(layoutUS, appointment.Date)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid appointment date format"})
	}
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Database.Db.Create(&appointment)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	return c.Status(201).JSON(result)
}
