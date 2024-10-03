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

	// Get environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	// Construct the DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Automatically migrate the schema for the Well model
	err = DB.AutoMigrate(&models.Ship{}, &models.Harbor{}, &models.User{}, &models.Ticket{}, &models.TicketClass{}, &models.Pathway{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully.")
}
