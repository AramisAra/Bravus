package sheetsapi

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func CreateSheet(c *fiber.Ctx) error {
	ctx := context.Background()
	name := c.Params("name")
	if name == "" {
		name = c.Query("name")
	}

	// Get the authorization code from the URL query parameters
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Authorization code not found")
	}

	// Exchange the authorization code for an access token
	tok, err := config.Exchange(ctx, code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Unable to exchange code for token: " + err.Error())
	}

	// Save the token to a file or use it to create a service client
	saveToken(tok)

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
	sheetid := c.Params("id")
	if sheetid == "" {
		sheetid = c.Query("id")
	}

	tok, err := ReadFromFile("token.json")
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
		c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(200).JSON(resp)
}

func GetSheetValues(c *fiber.Ctx) error {
	ctx := context.Background()
	sheetid := c.Params("id")
	if sheetid == "" {
		sheetid = c.Query("id")
	}

	tok, err := ReadFromFile("token.json")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": "No token file"})
	}

	client := config.Client(ctx, tok)
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}

	resp, err := srv.Spreadsheets.Values.BatchGet(sheetid).Ranges("A1:W337").Context(ctx).Do()
	if err != nil {
		c.Status(400).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.Status(200).JSON(resp)
}
