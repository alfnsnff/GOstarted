package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"GOstarted/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Local development using localhost PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", 
		os.Getenv("DB_HOST"), 
		os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}


	// Automatically migrate the schema for the Well model
	err = DB.AutoMigrate(&models.Ship{}, &models.Harbor{}, &models.User{}, &models.Ticket{}, &models.TicketClass{}, &models.Pathway{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully.")
}
