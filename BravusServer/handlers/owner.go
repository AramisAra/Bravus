package handlers

import (
	"time"

	"github.com/AramisAra/BravusServer/database"
	"github.com/AramisAra/BravusServer/database/models"
	"github.com/AramisAra/BravusServer/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/surrealdb/surrealdb.go"
)

func RegisterOwner(c *fiber.Ctx) error {

	db := database.ConnectDb()
	defer db.Close()

	var owner models.Owner
	createdAt := time.Now().Format(time.RFC3339)

	if err := c.BodyParser(&owner); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	hashedPassword, err := utils.HashPassword(owner.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Failed to crypt"})
	}

	owner.Password = hashedPassword

	query := "CREATE Owner CONTENT {createdAt: $createdAt, updateAt: $updateAt name: $name, phone: $phone, email: $email, password, $password, carrer: $carrer, appointments: $appointments}"
	params := map[string]interface{}{
		"createdAt":    createdAt,
		"updateAt":     createdAt,
		"name":         owner.Name,
		"phone":        owner.Phone,
		"email":        owner.Email,
		"password":     owner.Password,
		"carrer":       owner.Career,
		"appointments": owner.Appointments,
	}

	_, err = db.Query(query, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(owner)
}

func LoginOwner(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	db := database.ConnectDb()
	defer db.Close()

	query := "SELECT * FROM Owner WHERE email: $email"
	params := map[string]interface{}{
		"email": input.Email,
	}

	response, err := db.Query(query, params)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var results []map[string]interface{}
	if err := surrealdb.Unmarshal(response, &results); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	if len(results) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	// Extract the password from the nested structure
	data, ok := results[0]["result"].([]interface{})
	if !ok || len(data) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}

	ownerData, ok := data[0].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "password retrieval error"})
	}

	storedPassword, ok := ownerData["password"].(string)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "password retrieval error"})
	}

	if err := utils.ComparePasswords(storedPassword, input.Password); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}
}
