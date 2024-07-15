package database

import (
	"log"
	"os"

	dbmodels "github.com/AramisAra/BravusBackend/database/dbmodels"
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
	// Set logger mode to Info
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// Perform migrations
	db.AutoMigrate(&dbmodels.Client{}, &dbmodels.Owner{}, &dbmodels.Animal{},
		&dbmodels.Service{}, &dbmodels.Appointment{})

	Database = DbInstance{Db: db}
}
