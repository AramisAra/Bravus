package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

// Creates Service
func CreateService(c *fiber.Ctx) error {
	// Var == to link param UUID
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	// Checks if its a uuid
	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	// database query for the owner how making the service
	var owner dbmodels.Owner

	database.Database.Db.Find(&owner, "id = ?", id)

	var service dbmodels.Service

	// Start of creating the service
	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	// The Creates the service with the owner id foreign key
	service.OwnerID = owner.ID
	database.Database.Db.Create(&service)
	responseService := dbmodels.CreateServiceResponse(service)

	return c.Status(200).JSON(responseService)
}

// List all the service alone
func ListService(c *fiber.Ctx) error {
	var services []dbmodels.Service

	database.Database.Db.Find(&services)
	responseService := []dbmodels.ServiceSerializer{}

	for _, service := range services {
		responseService = append(responseService, dbmodels.CreateServiceResponse(service))
	}

	return c.Status(200).JSON(responseService)
}

// Updates service
func UpdateService(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

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

// Delete service
func DeleteService(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

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
