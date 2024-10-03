package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"GOstarted/models"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect() {
	var err error

	// Load .env file only for local development (if necessary)
	if os.Getenv("ENV") == "development" {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Prepare the DSN based on environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // for Cloud SQL, this should be the socket path
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	// Open a database connection using Gorm
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Automatically migrate the schema for the defined models
	err = DB.AutoMigrate(&models.Ship{}, &models.Harbor{}, &models.User{}, &models.Ticket{}, &models.TicketClass{}, &models.Pathway{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully.")
}
