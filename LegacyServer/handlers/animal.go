package handlers

import (
	database "github.com/AramisAra/BravusBackend/database"
	"github.com/AramisAra/BravusBackend/database/dbmodels"
	"github.com/gofiber/fiber/v2"
)

// CreateAnimal is a handler function that creates a new animal record in the database.
// It expects a valid UUID as a query parameter and a JSON payload containing the animal data in the request body.
// If the UUID is invalid, it returns a 400 Bad Request response.
// If the animal data cannot be parsed from the request body, it returns a 400 Bad Request response with the error message.
// If there is an error while creating the animal record in the database, it returns a 500 Internal Server Error response with the error message.
// Otherwise, it returns a 201 Created response with the created animal record in the JSON format.
func CreateAnimal(c *fiber.Ctx) error {
	id := c.Query("uuid")
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

// UpdateAnimal updates an animal record in the database based on the provided UUID.
// It expects the UUID to be either passed as a parameter or as a query parameter.
// If the UUID is invalid, it returns a Bad Request response.
// The function retrieves the animal record from the database using the UUID.
// It then parses the request body to get the updated animal information.
// The function updates the animal record with the new information and saves it back to the database.
// Finally, it returns a JSON response with the updated animal record.
func UpdateAnimal(c *fiber.Ctx) error {
	id := c.Query("uuid")
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

// DeleteAnimal deletes an animal from the database based on the provided UUID.
// It returns an error if the UUID is invalid or if there was an issue deleting the animal.
// If the animal is successfully deleted, it returns a JSON response indicating success.
func DeleteAnimal(c *fiber.Ctx) error {
	id := c.Query("uuid")
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
