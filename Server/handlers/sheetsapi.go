package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func GetApi() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the Env: ", err)
	}
	var API_KEY = os.Getenv("API_KEY")
	return API_KEY
}

func GetTable(c *fiber.Ctx) error {
	API_KEY := GetApi()
	userSheetName := c.Params("sheet")
	response, err := http.Get(API_KEY + "?sheet=" + userSheetName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(data)
}

func GetSheets(c *fiber.Ctx) error {
	API_KEY := GetApi()
	response, err := http.Get(API_KEY + "/sheets")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var data interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(data)
}
