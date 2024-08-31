package handlers

import (
	"time"

	"github.com/AramisAra/BravusServer/database"
	"github.com/AramisAra/BravusServer/database/models"
	"github.com/gofiber/fiber/v2"
)

func CreateService(c *fiber.Ctx) error {

	Db := database.ConnectDb()
	defer Db.Close()

	createdAt := time.Now().Format(time.RFC3339)
	service := models.Service{}

	err := c.BodyParser(&service)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	query := "CREATE Service CONTENT { createdAt: $createdAt, NameService: $NameService, ServiceDesc: $ServiceDesc, Price: $Price, OwnerID: $OwnerID}"
	values := map[string]interface{}{
		"createdAt":   createdAt,
		"NameService": service.NameService,
		"ServiceDesc": service.ServiceDesc,
		"Price":       service.Price,
		"OwnerID":     service.OwnerID,
	}

	_, err = Db.Query(query, values)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to create service"})
	}

	return c.Status(200).JSON(service)
}
