package handlers

import (
	"time"

	"github.com/AramisAra/BravusServer/config"
	"github.com/AramisAra/BravusServer/database"
	"github.com/AramisAra/BravusServer/database/models"
	"github.com/AramisAra/BravusServer/utils"
	"github.com/dgrijalva/jwt-go"
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

	// Check if the email is already in use
	emailCheckQuery := `SELECT * FROM Owner, Owner WHERE email = $email`
	emailCheckValues := map[string]interface{}{
		"email": owner.Email,
	}

	emailCheckResponse, err := db.Query(emailCheckQuery, emailCheckValues)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Assuming response is a JSON array, unmarshal into a slice of maps
	var emailCheckResult []map[string]interface{}
	if err := surrealdb.Unmarshal(emailCheckResponse, &emailCheckResult); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	resultData, ok := emailCheckResult[0]["result"].([]interface{})
	if !ok || len(resultData) == 0 {
		hashedPassword, err := utils.HashPassword(owner.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Error": "Failed to crypt"})
		}

		owner.Password = hashedPassword

		query := "CREATE Owner CONTENT {createdAt: $createdAt, updateAt: $updateAt, name: $name, phone: $phone, email: $email, password: $password, career: $career, appointments: $appointments}"
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
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email already exist"})
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

	ownerid := ownerData["id"].(string)
	owneremail := ownerData["email"].(string)
	claims := jwt.MapClaims{
		"ID":    ownerid,
		"email": owneremail,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not login"})
	}

	return c.JSON(models.LoginResponse{ID: ownerid, Token: t})
}

// ListOwner handles the request to list all Owners.
func ListOwner(c *fiber.Ctx) error {
	db := database.ConnectDb()
	defer db.Close()

	response, err := db.Query("SELECT * FROM Owner", map[string]interface{}{
		"tb": "Owner",
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error})
	}

	return c.Status(200).JSON(response)
}

func GetOwner(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No ID"})
	} else {
		checker := utils.IsValidOwnerString(id)
		if !checker {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid id"})
		}
	}

	db := database.ConnectDb()
	defer db.Close()

	// Queries for the Owner
	response, err := db.Query("SELECT * From Owner WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(response)
}

// UpdateOwner updates the Owner information in the database based on the provided ID..
func UpdateOwner(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No ID"})
	} else {
		checker := utils.IsValidOwnerString(id)
		if !checker {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid id"})
		}
	}

	db := database.ConnectDb()
	defer db.Close()

	// Fetch the existing Owner data from the database
	query := `SELECT * FROM Owner WHERE id = $id`
	values := map[string]interface{}{
		"id": id,
	}

	response, err := db.Query(query, values)
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
	existingOwnerData, ok := existingData[0].(map[string]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid Owner data format"})
	}

	// Parse the request body into the Owner struct
	var owner models.Owner
	if err := c.BodyParser(&owner); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": err.Error()})
	}

	// Only update the fields that are provided in the request body
	if owner.Name != "" {
		existingOwnerData["name"] = owner.Name
	}
	if owner.Email != "" {
		existingOwnerData["email"] = owner.Email
	}
	if owner.Phone != "" {
		existingOwnerData["phone"] = owner.Phone
	}
	if owner.Career != "" {
		existingOwnerData["career"] = owner.Career
	}
	if owner.Services != nil {
		existingOwnerData["services"] = owner.Services
	}
	if owner.Appointments != nil {
		existingOwnerData["appointments"] = owner.Appointments
	}

	// Update the updateAt field
	updateAt := time.Now().Format(time.RFC3339)

	// Build the update query
	updateQuery := "UPDATE $id CONTENT {name: $name, email: $email, phone: $phone, password: $password, createdAt: $createdAt, updateAt: $updateAt, career: $career, services: $services, appointments: $appointments}"
	updateValues := map[string]interface{}{
		"id":           id,
		"createdAt":    existingOwnerData["createdAt"],
		"updateAt":     updateAt,
		"name":         existingOwnerData["name"],
		"phone":        existingOwnerData["phone"],
		"email":        existingOwnerData["email"],
		"password":     existingOwnerData["password"],
		"career":       existingOwnerData["career"],
		"services":     existingOwnerData["services"],
		"appointments": existingOwnerData["appointments"],
	}

	// Execute the update query
	if _, err := db.Query(updateQuery, updateValues); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(existingOwnerData)
}

func DeleteOwner(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No ID"})
	} else {
		checker := utils.IsValidOwnerString(id)
		if !checker {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not a valid id"})
		}
	}

	db := database.ConnectDb()
	defer db.Close()

	response, err := db.Query("DELETE FROM Owner WHERE id = $id;", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(response)
}
