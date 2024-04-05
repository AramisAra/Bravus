package database

import (
	"log"
	"os"

	models "github.com/AramisAra/Grooming_App/Database/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

// ConnectDb establishes a connection to the database and performs necessary migrations.
func ConnectDb() {
	sqlDb := "host=devdatabase.c582ws226em9.us-east-1.rds.amazonaws.com user=Ara_Bard password=DogGrooming port=5432 database=dog_grooming"
	db, err := gorm.Open(postgres.Open(sqlDb), &gorm.Config{})
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
