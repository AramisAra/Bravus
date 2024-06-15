package routes

import (
	database "github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func isValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// Routes Functions For Clients With Animals
func CreateClientAndAnimal(c *fiber.Ctx) error {
	// Creates Client first
	var client models.Client

	if err := c.BodyParser(&client); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Database.Db.Create(&client)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	// Get the newly created client ID
	clientID := client.ID

	// Creates Animals second
	var animal models.Animals

	animal.Client_id = clientID // Set foreign key to client ID
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// Insert animal with the foreign key set
	result = database.Database.Db.Create(&animal)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	Response := database.CreateJoinResultClient(client, animal)

	return c.Status(201).JSON(Response)
}

// For Creating A Second Animals
func CreateAnimal(c *fiber.Ctx) error {
	// Creates Client first
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	client := models.Client{}

	database.Database.Db.Find(&client, "id = ?", id)

	ClientID := client.ID

	var animal models.Animals

	animal.Client_id = ClientID
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Database.Db.Create(&animal)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	Response := database.CreateAnimalResponse(animal)

	return c.Status(201).JSON(Response)
}

// Get Clients with Animals
func GetClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	client := models.Client{}

	database.Database.Db.Preload("Animals").Find(&client, "id = ?", id)

	response := database.CreateClientResponse(client)

	return c.Status(200).JSON(response)
}

func ListClients(c *fiber.Ctx) error {
	clients := []models.Client{}

	database.Database.Db.Preload("Animals").Find(&clients)
	responseClients := []database.ClientSerializer{}

	for _, client := range clients {
		responseClients = append(responseClients, database.CreateClientResponse(client))
	}

	return c.Status(200).JSON(responseClients)
}

func UpdateClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	client := models.Client{}

	database.Database.Db.Joins("Animals").Find(&client, "id = ?", id)

	var updateClient database.UpdateClientInput

	if err := c.BodyParser(&updateClient); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	client.Full_Name = updateClient.FullName
	client.Email = updateClient.Email
	client.Phone = updateClient.Phone

	database.Database.Db.Omit("Animals").Save(&client)

	responseClient := database.CreateClientResponse(client)
	return c.Status(200).JSON(responseClient)
}

func DeleteClient(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	client := models.Client{}

	database.Database.Db.Find(&client)
	if err := database.Database.Db.Delete(&client).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Client was deleted")
}

func ListAnimals(c *fiber.Ctx) error {
	animals := []models.Animals{}

	database.Database.Db.Find(&animals)
	responseAnimal := []database.AnimalSerializer{}

	for _, animal := range animals {
		responseAnimal = append(responseAnimal, database.CreateAnimalResponse(animal))
	}

	return c.Status(200).JSON(responseAnimal)
}

func GetAnimal(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	animal := models.Animals{}

	database.Database.Db.Find(&animal)
	responseAnimal := database.CreateAnimalResponse(animal)

	return c.Status(200).JSON(responseAnimal)
}

func UpdateAnimal(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	animal := models.Animals{}

	database.Database.Db.Find(&animal, "id = ?", id)

	var updateAnimal database.UpdateAnimalInput

	if err := c.BodyParser(&updateAnimal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	animal.Animal_Name = updateAnimal.Animal_Name
	animal.Animal_Specie = updateAnimal.Animal_Specie
	animal.Animal_Age = updateAnimal.Animal_Age

	database.Database.Db.Save(&animal)

	responseAnimal := database.CreateAnimalResponse(animal)
	return c.Status(200).JSON(responseAnimal)
}

func GetAppointment(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	var client models.Client

	database.Database.Db.Preload("Appointments").Preload("Animals").Find(&client)

	response := database.CreateClientResponse(client)

	return c.Status(200).JSON(response)
}
