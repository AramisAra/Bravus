package sheetsapi

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

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

func Start() {
	ClientGetter()
}

func ClientGetter() {
	// Load The Secrets Files
	b, err := os.ReadFile(clientSecretFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Parse the client secret file to config
	config, err = google.ConfigFromJSON(b, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
}

func ReadFromFile(filename string) (*oauth2.Token, error) {
	path := filepath.Join(".tokens", filename)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(token *oauth2.Token, filename string) {
	path := filepath.Join(".tokens", filename+".json")
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("Unable to create token file: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func AuthGoogle(c *fiber.Ctx) error {
	id := c.Query("uuid")

	if !isValidUUID(id) {
		println(id)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid UUID")
	}
	// Generate the URL for the authorization request.
	ClientGetter()
	authURL := config.AuthCodeURL(id, oauth2.AccessTypeOffline)
	return c.Redirect(authURL)
}
