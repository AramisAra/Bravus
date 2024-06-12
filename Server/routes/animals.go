package routes

import (
	database "github.com/AramisAra/GroomingApp/database"
	models "github.com/AramisAra/GroomingApp/models"
	"github.com/AramisAra/GroomingApp/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateAnimal(c *fiber.Ctx) error {
	var animal models.Animals

	if err := c.BodyParser(animal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&animal)
	responseAnimal := utils.CreateAnimalResponse(animal)

	return c.Status(200).JSON(responseAnimal)
}

func ListAnimals(c *fiber.Ctx) error {
	animals := []models.Animals{}

	database.Database.Db.Find(&animals)
	responseAnimal := []utils.AnimalSerializer{}

	for _, animal := range animals {
		responseAnimal = append(responseAnimal, utils.CreateAnimalResponse(animal))
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

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	animal := models.Animals{}

	if err := utils.FindAnimal(parsedId, &animal); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	responseAnimal := utils.CreateAnimalResponse(animal)

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

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	animal := models.Animals{}
	if err := utils.FindAnimal(parsedId, &animal); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	var updateAnimal utils.UpdateAnimalInput

	if err := c.BodyParser(&updateAnimal); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	animal.Animal_Name = updateAnimal.Animal_Name
	animal.Animal_Specie = updateAnimal.Animal_Specie
	animal.Animal_Age = updateAnimal.Animal_Age

	responseAnimal := utils.CreateAnimalResponse(animal)
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

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error parsing ID: " + err.Error())
	}

	animal := models.Animals{}
	if err := utils.FindAnimal(parsedId, &animal); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&animal).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON("Animal Was Deleted")
}
