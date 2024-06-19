package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func CreateService(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	var owner dbmodels.Owner

	database.Database.Db.Find(&owner, "id = ?", id)

	var service dbmodels.Service

	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	service.OwnerID = owner.ID
	database.Database.Db.Create(&service)
	responseService := dbmodels.CreateServiceResponse(service)

	return c.Status(200).JSON(responseService)
}

func ListService(c *fiber.Ctx) error {
	var services []dbmodels.Service

	database.Database.Db.Find(&services)
	responseService := []dbmodels.ServiceSerializer{}

	for _, service := range services {
		responseService = append(responseService, dbmodels.CreateServiceResponse(service))
	}

	return c.Status(200).JSON(responseService)
}

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
