package routes

import (
	database "github.com/AramisAra/GroomingApp/database"
	"github.com/AramisAra/GroomingApp/models"
	"github.com/AramisAra/GroomingApp/utils"
	"github.com/gofiber/fiber/v2"
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
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please enter valid id: An integer")
	}

	service := models.Service{}

	if err := utils.FindService(id, &service); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseService := utils.CreateServiceResponse(service)

	return c.Status(200).JSON(responseService)
}

func UpdateService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please enter valid id: An integer")
	}
	service := models.Service{}
	if err := utils.FindService(id, &service); err != nil {
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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please enter valid id: An integer")
	}
	service := models.Service{}
	if err := utils.FindService(id, &service); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&service).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Service Was Deleted")
}
