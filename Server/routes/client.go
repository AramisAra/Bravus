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
		return c.Status(400).JSON("Please enter a valid id: An integer")
	}

	client := models.Client{}

	if err := utils.FindClient(id, &client); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseClient := utils.CreateClientResponse(client)

	return c.Status(200).JSON(responseClient)
}

func UpdateClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please enter a valid id: An integer")
	}
	client := models.Client{}
	if err := utils.FindClient(id, &client); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateClient utils.UpdateClientInput

	if err := c.BodyParser(&updateClient); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	client.Full_Name = updateClient.FullName
	client.Email = updateClient.Email
	client.Phone = updateClient.Phone
	database.Database.Db.Save(&client)

	responseClient := utils.CreateClientResponse(client)
	return c.Status(200).JSON(responseClient)
}

func DeleteClient(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please enter a valid id: An integer")
	}
	client := models.Client{}

	if err := utils.FindClient(id, &client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&client).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Client was deleted")
}
