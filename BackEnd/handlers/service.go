package handlers

import (
	"fmt"

	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

// CreateService is a handler function that creates a new service.
// It expects a valid UUID as a query parameter and a JSON payload representing the service.
// The function retrieves the owner associated with the UUID from the database,
// parses the JSON payload into a service object, assigns the owner ID to the service,
// and creates the service in the database.
// Finally, it returns a JSON response containing the created service.
func CreateService(c *fiber.Ctx) error {
	id := c.Query("uuid")
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var owner dbmodels.Owner

	database.Database.Db.Find(&owner, "id = ?", id)

	var service dbmodels.Service

	rawBody := c.Body()
	fmt.Println("Raw request body:", string(rawBody))

	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	service.OwnerID = owner.ID
	database.Database.Db.Create(&service)
	responseService := dbmodels.CreateServiceResponse(service)

	return c.Status(200).JSON(responseService)
}

// ListService retrieves a list of services from the database and returns them as a JSON response.
// It uses the fiber.Ctx parameter to handle the HTTP request and response.
// The function queries the database for all services, serializes them using the ServiceSerializer,
// and returns the serialized services as a JSON response.
// The function returns an error if there is an issue with the database or the JSON serialization.
// FILEPATH: /home/ara/Bravus/BackEnd/handlers/service.go
func ListService(c *fiber.Ctx) error {
	var services []dbmodels.Service

	database.Database.Db.Find(&services)
	responseService := []dbmodels.ServiceSerializer{}

	for _, service := range services {
		responseService = append(responseService, dbmodels.CreateServiceResponse(service))
	}

	return c.Status(200).JSON(responseService)
}

// UpdateService updates a service based on the provided UUID.
// It retrieves the service from the database, updates its properties with the values from the request body,
// and saves the updated service back to the database. Finally, it returns the updated service as a JSON response.
func UpdateService(c *fiber.Ctx) error {
	id := c.Query("uuid")
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}

	var service dbmodels.Service

	database.Database.Db.Find(&service, "id = ?", id)

	var UpdateService dbmodels.UpdateServiceInput

	if err := c.BodyParser(&UpdateService); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	service.NameService = UpdateService.ServiceName
	service.ServiceDesc = UpdateService.ServiceDesc
	service.Price = UpdateService.Price
	service.ServiceCode = UpdateService.ServiceCode

	database.Database.Db.Save(&service)

	responseService := dbmodels.CreateServiceResponse(service)
	return c.Status(200).JSON(responseService)
}

// DeleteService deletes a service from the database based on the provided UUID.
// It returns an error if the UUID is invalid or if the service cannot be deleted.
func DeleteService(c *fiber.Ctx) error {
	id := c.Query("uuid")

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}

	var service dbmodels.Service

	database.Database.Db.Find(&service, "id = ?", id)
	if err := database.Database.Db.Delete(&service); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "unable to delete service"})
	}

	return c.Status(200).JSON("Service was deleted")
}
