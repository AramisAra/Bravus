package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

func CreateService(c *fiber.Ctx) error {
	var service dbmodels.Service

	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&service)
	responseService := dbmodels.CreateServiceResponse(service)

	return c.Status(200).JSON(responseService)
}
