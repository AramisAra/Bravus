package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// Will load ".env" file into the server
func GetDot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the Env: ", err)
	}
}
