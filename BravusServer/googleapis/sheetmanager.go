package googleapis

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func isValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func CreateSheet(c *fiber.Ctx) error {
	ctx := context.Background()
	name := c.Query("name")
	uuid := c.Query("uuid")

	// Read token from the json
	tok, err := ReadFromFile(uuid + ".json")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "No token file"})
	}

	// Create a new Sheets service client
	client := config.Client(ctx, tok)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to create Sheets service: " + err.Error())
	}

	// Use the service client to create a new spreadsheet
	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: name,
		},
	}
	resp, err := srv.Spreadsheets.Create(spreadsheet).Context(ctx).Do()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to create spreadsheet: " + err.Error())
	}

	return c.Status(201).JSON(resp)
}

func GetSheet(c *fiber.Ctx) error {
	ctx := context.Background()
	sheetid := c.Query("id")
	uuid := c.Query("uuid")

	tok, err := ReadFromFile(uuid + ".json")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "No token file"})
	}

	client := config.Client(ctx, tok)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}

	resp, err := srv.Spreadsheets.Get(sheetid).Context(ctx).Do()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}

	response := fiber.Map{
		"Title":         resp.Properties.Title,
		"spreadsheetId": resp.SpreadsheetId,
	}

	return c.Status(200).JSON(response)
}

func GetSheetValues(c *fiber.Ctx) error {
	ctx := context.Background()
	sheetid := c.Query("id")
	uuid := c.Query("uuid")

	tok, err := ReadFromFile(uuid + ".json")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "No token file"})
	}

	client := config.Client(ctx, tok)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}

	resp, err := srv.Spreadsheets.Values.BatchGet(sheetid).Ranges("A1:O100").Context(ctx).Do()
	if err != nil {
		c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(200).JSON(resp)
}
