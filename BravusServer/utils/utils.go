package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func GetDot(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the Env: ", err)
	}
}
