package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

// CreateAppointment is a handler function that creates a new appointment.
// It takes a fiber.Ctx object as a parameter and returns an error.
// The function retrieves the user and owner IDs from the query parameters,
// validates them as UUIDs, and then proceeds to retrieve the corresponding
// client and owner records from the database. It then parses the request
// body into an appointment object and saves it to the database. Finally,
// it constructs a response object and returns it as JSON.
func CreateAppointment(c *fiber.Ctx) error {
	userid := c.Query("iduser")
	if !isValidUUID(userid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	ownerid := c.Query("idowner")
	if !isValidUUID(ownerid) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	// Data Management
	var client dbmodels.Client
	var owner dbmodels.Owner

	database.Database.Db.Find(&client, "id = ?", userid)
	database.Database.Db.Find(&owner, "id = ?", ownerid)

	var appointment dbmodels.Appointment

	appointment.ClientID = client.ID
	appointment.OwnerID = owner.ID

	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Database.Db.Create(&appointment)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}
	// response Handling
	response := dbmodels.CreateAppointmentResponse(appointment)

	return c.Status(201).JSON(response)
}

// DeleteAppointment deletes an appointment from the database based on the provided UUID.
// It returns an error if the UUID is invalid or if there was an issue deleting the appointment.
// If the appointment is successfully deleted, it returns a JSON response with a success message.
func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Query("uuid")
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	appointment := dbmodels.Appointment{}

	database.Database.Db.Find(&appointment, "id = ?", id)
	if err := database.Database.Db.Delete(&appointment).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Appointment was deleted")
}

// UpdateAppointment updates an appointment based on the provided UUID.
// It retrieves the appointment from the database using the UUID, then updates the appointment's date and time
// with the values provided in the request body. Finally, it saves the updated appointment to the database
// and returns the updated appointment as a JSON response.
// Parameters:
// - c: The fiber.Ctx object representing the HTTP request context.
// Returns:
// - An error if there was an issue with the request or database operation, otherwise returns nil.
func UpdateAppointment(c *fiber.Ctx) error {
	id := c.Query("uuid")

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var appointment dbmodels.Appointment
	database.Database.Db.Find(&appointment, "id = ?", id)

	var updateAppointment dbmodels.AppointmentUpdater

	if err := c.BodyParser(&updateAppointment); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	appointment.Date = updateAppointment.Date
	appointment.Time = updateAppointment.Time

	database.Database.Db.Save(&appointment)

	responseAppointment := dbmodels.CreateAppointmentResponse(appointment)
	return c.Status(200).JSON(responseAppointment)
}
