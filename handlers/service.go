package handlers

import (
	"time"

	"github.com/AramisAra/BravusServer/database"
	"github.com/AramisAra/BravusServer/database/models"
	"github.com/AramisAra/BravusServer/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
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

	query := "CREATE Service CONTENT { createdAt: $createdAt, updateAt: $updateAt, NameService: $NameService, ServiceDesc: $ServiceDesc, Price: $Price, OwnerID: $OwnerID}"
	values := map[string]interface{}{
		"createdAt":   createdAt,
		"updateAt":    createdAt,
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

func ListService(c *fiber.Ctx) error {

	Db := database.ConnectDb()
	defer Db.Close()

	query := "SELECT * FROM Service"
	value := map[string]interface{}{
		"tb": "Service",
	}

	response, err := Db.Query(query, value)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(response)
}

func UpdateService(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No ID"})
	} else {
		checker := utils.IsValidServiceString(id)
		if !checker {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid id"})
		}
	}

	Db := database.ConnectDb()
	defer Db.Close()

	query := `SELECT * FROM Service WHERE id = $id`
	values := map[string]interface{}{
		"id": id,
	}

	response, err := Db.Query(query, values)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Assuming response is a JSON array, unmarshal into a slice of maps
	var results []map[string]interface{}
	if err := surrealdb.Unmarshal(response, &results); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Extract the existing Owner data from the nested structure
	existingData, ok := results[0]["result"].([]interface{})
	if !ok || len(existingData) == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid Owner data format"})
	}
	existingServiceData, ok := existingData[0].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid Owner data format"})
	}

	// Parse the request body into the Owner struct
	var service models.Service
	if err := c.BodyParser(&service); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}

	// Only update the fields that are provided in the request body
	if service.NameService != "" {
		existingServiceData["NameService"] = service.NameService
	}
	if service.ServiceDesc != "" {
		existingServiceData["ServiceDesc"] = service.ServiceDesc
	}
	if service.Price < 0 {
		existingServiceData["Price"] = service.Price
	}

	// Update the updateAt field
	updateAt := time.Now().Format(time.RFC3339)

	// Build the update query
	updateQuery := "UPDATE $id CONTENT {NameService: $NameService, ServiceDesc: $ServiceDesc, OwnerID: $OwnerID, Price: $Price, createdAt: $createdAt, updateAt: $updateAt}"
	updateValues := map[string]interface{}{
		"id":          id,
		"createdAt":   existingServiceData["createdAt"],
		"updateAt":    updateAt,
		"NameService": existingServiceData["NameService"],
		"ServiceDesc": existingServiceData["ServiceDesc"],
		"Price":       existingServiceData["Price"],
		"OwnerID":     existingServiceData["OwnerID"],
	}

	// Execute the update query
	if _, err := Db.Query(updateQuery, updateValues); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(existingServiceData)
}

// This delete owner base on the give ID
func DeleteService(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No ID"})
	} else {
		checker := utils.IsValidServiceString(id)
		if !checker {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid id"})
		}
	}

	db := database.ConnectDb()
	defer db.Close()

	response, err := db.Query("DELETE FROM Service WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(response)
}
