package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

var (
	// The path to the client secret file downloaded from Google Cloud Console
	clientSecretFile = "client_secret.json"

	// OAuth2 config
	config *oauth2.Config
)

func main() {
	// Load the client secret file
	b, err := os.ReadFile(clientSecretFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Parse the client secret file to config
	config, err = google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// Generate the URL for the authorization request.
		authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
		return c.SendString(fmt.Sprintf("Go to the following link in your browser: \n%v\n", authURL))
	})

	app.Get("/callback", func(c *fiber.Ctx) error {
		ctx := context.Background()

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
		srv, err := sheets.New(client)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Unable to create Sheets service: " + err.Error())
		}

		// Use the service client to create a new spreadsheet
		spreadsheet := &sheets.Spreadsheet{
			Properties: &sheets.SpreadsheetProperties{
				Title: "My New Spreadsheet",
			},
		}
		resp, err := srv.Spreadsheets.Create(spreadsheet).Context(ctx).Do()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Unable to create spreadsheet: " + err.Error())
		}

		return c.SendString(fmt.Sprintf("Spreadsheet ID: %s\n", resp.SpreadsheetId))
	})

	log.Fatal(app.Listen(":8080"))
}

func saveToken(token *oauth2.Token) {
	f, err := os.Create("token.json")
	if err != nil {
		log.Fatalf("Unable to create token file: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
