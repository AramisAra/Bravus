package sheetsapi

import (
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

func ReadFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveToken(token *oauth2.Token) {
	f, err := os.Create("token.json")
	if err != nil {
		log.Fatalf("Unable to create token file: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func AuthGoogle(c *fiber.Ctx) error {
	// Generate the URL for the authorization request.
	ClientGetter()
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.SendString(fmt.Sprintf("Go to the following link in your browser: \n%v\n", authURL))
}
