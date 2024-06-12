package routes

import (
	database "github.com/AramisAra/GroomingApp/database"
	"github.com/AramisAra/GroomingApp/models"
	"github.com/AramisAra/GroomingApp/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateService(c *fiber.Ctx) error {
	var service models.Service

	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&service)
	responseService := utils.CreateServiceResponse(service)

	return c.Status(200).JSON(responseService)
}

func ListService(c *fiber.Ctx) error {
	services := []models.Service{}

	database.Database.Db.Find(&services)
	responseService := []utils.ServiceSerializer{}

	for _, service := range services {
		responseService = append(responseService, utils.CreateServiceResponse(service))
	}

	return c.Status(200).JSON(responseService)
}

func GetService(c *fiber.Ctx) error {
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

	service := models.Service{}

	if err := utils.FindService(parsedId, &service); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseService := utils.CreateServiceResponse(service)

	return c.Status(200).JSON(responseService)
}

func UpdateService(c *fiber.Ctx) error {
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

	service := models.Service{}
	if err := utils.FindService(parsedId, &service); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateService utils.UpdateServiceInput

	if err := c.BodyParser(&updateService); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	service.NameService = updateService.ServiceName
	service.ServiceDesc = updateService.ServiceDesc
	service.ServiceCode = updateService.ServiceCode

	responseService := utils.CreateServiceResponse(service)
	return c.Status(200).JSON(responseService)
}

func DeleteService(c *fiber.Ctx) error {
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

	service := models.Service{}
	if err := utils.FindService(parsedId, &service); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&service).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Service Was Deleted")
}
