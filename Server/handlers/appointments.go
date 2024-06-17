package handlers

import (
	database "github.com/AramisAra/GroomingApp/database"
	"github.com/AramisAra/GroomingApp/models"
	"github.com/gofiber/fiber/v2"
)

func CreateAppointment(c *fiber.Ctx) error {
	// UUID Valid Checking
	userid := c.Params("uuidUser")
	if userid == "" {
		userid = c.Query("uuidUser")
	}
	if !isValidUUID(userid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	ownerid := c.Params("uuidOwner")
	if ownerid == "" {
		ownerid = c.Query("uuidOwner")
	}
	if !isValidUUID(ownerid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	animalid := c.Params("uuidAnimal")
	if animalid == "" {
		animalid = c.Query("uuidAnimal")
	}
	if !isValidUUID(animalid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	// Data Management
	var client models.Client
	var owner models.Owner
	var animal models.Animals

	database.Database.Db.Find(&client, "id = ?", userid)
	database.Database.Db.Find(&owner, "id = ?", ownerid)
	database.Database.Db.Find(&animal, "id = ?", animalid)

	var appointment models.Appointment

	appointment.ClientID = client.ID
	appointment.OwnerID = owner.ID
	appointment.AnimalID = animal.ID

	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Database.Db.Create(&appointment)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}
	// response Handling
	response := database.CreateAppointmentResponse(&appointment)

	return c.Status(201).JSON(response)
}

func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	appointment := models.Appointment{}

	database.Database.Db.Find(&appointment)
	if err := database.Database.Db.Delete(&appointment).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Appointment was deleted")
}

func UpdateAppointment(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var appointment models.Appointment
	database.Database.Db.Find(&appointment, "id = ?", id)

	var updateAppointment database.AppointmentUpdater

	if err := c.BodyParser(&updateAppointment); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	appointment.Date = updateAppointment.Date
	appointment.Time = updateAppointment.Time

	database.Database.Db.Save(&appointment)

	responseAppointment := database.CreateAppointmentResponse(&appointment)
	return c.Status(200).JSON(responseAppointment)
}
