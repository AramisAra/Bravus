package database

import (
	"log"
	"os"

	models "github.com/AramisAra/GroomingApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

// ConnectDb establishes a connection to the database and performs necessary migrations.
func ConnectDb(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect: ", err.Error())
		os.Exit(2)
	}
	log.Println("Connected")

	// Logs Creations
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running Migrations")
	// migrations
	db.AutoMigrate(&models.Client{}, &models.Appointment{}, &models.Service{})

	Database = DbInstance{Db: db}
}
