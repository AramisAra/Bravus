package routes

import (
	database "github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	"github.com/AramisAra/GroomingApp/utils"
	"github.com/gofiber/fiber/v2"
)

func CreateClient(c *fiber.Ctx) error {
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&client)
	responseClient := utils.CreateClientResponse(client)

	return c.Status(200).JSON(responseClient)
}

func ListClients(c *fiber.Ctx) error {
	clients := []models.Client{}

	database.Database.Db.Find(&clients)
	responseClients := []utils.ClientSerializer{}

	for _, client := range clients {
		responseClients = append(responseClients, utils.CreateClientResponse(client))
	}

	return c.Status(200).JSON(responseClients)
}

func GetClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please enter a valid ID: It must be a number")
	}

	client := models.Client{}

	if err := utils.FindClient(id, &client); if err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseClient := utils.CreateClientResponse(client)

	return c.Status(200).JSON(responseClient)
}
