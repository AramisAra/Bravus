package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

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
	client := dbmodels.Client{}

	database.Database.Db.Find(&client, "id = ?", id)

	ClientID := client.ID

	var animal dbmodels.Animal

	animal.Client_id = ClientID
	if err := c.BodyParser(&animal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result := database.Database.Db.Create(&animal)

	if result.Error != nil {
		return c.Status(500).JSON(result.Error)
	}

	Response := dbmodels.CreateAnimalResponse(animal)

	return c.Status(201).JSON(Response)
}

func UpdateAnimal(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	animal := dbmodels.Animal{}

	database.Database.Db.Find(&animal, "id = ?", id)

	var updateAnimal dbmodels.UpdateAnimalInput

	if err := c.BodyParser(&updateAnimal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	animal.Animal_Name = updateAnimal.Animal_Name
	animal.Animal_Specie = updateAnimal.Animal_Specie
	animal.Animal_Age = updateAnimal.Animal_Age

	database.Database.Db.Save(&animal)

	responseAnimal := dbmodels.CreateAnimalResponse(animal)
	return c.Status(200).JSON(responseAnimal)
}

func DeleteAnimal(c *fiber.Ctx) error {
	id := c.Params("uuid")
	if id == "" {
		id = c.Query("uuid")
	}

	if !isValidUUID(id) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}

	animal := dbmodels.Animal{}

	database.Database.Db.Find(&animal, "id = ?", id)
	if err := database.Database.Db.Delete(&animal).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Animal was deleted")
}

/*
func ListAnimals(c *fiber.Ctx) error {
	animals := []dbmodels.Animal{}

	database.Database.Db.Find(&animals)
	responseAnimal := []dbmodels.AnimalSerializer{}

	for _, animal := range animals {
		responseAnimal = append(responseAnimal, dbmodels.CreateAnimalResponse(animal))
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

	animal := dbmodels.Animal{}

	database.Database.Db.Find(&animal)
	responseAnimal := dbmodels.CreateAnimalResponse(animal)

	return c.Status(200).JSON(responseAnimal)
}
*/
