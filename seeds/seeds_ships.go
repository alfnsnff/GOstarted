package seeds

import (
	"time"
    "log"
    "GOstarted/database"
    "GOstarted/models"
)

func SeedShips() {
    // Define some dummy well data
    ships := []models.Ship{
        {
			Name: "Aceh Hebat 1",    // Name of the ship
			Capacity:  100, // Passenger capacity
            // Set the geom field with a valid POINT geometry, e.g., `Geom: "POINT(lng lat)"`
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
        {
			Name: "Teluk Sinabang",    // Name of the ship
			Capacity:  100, // Passenger capacity
            // Set the geom field with a valid POINT geometry, e.g., `Geom: "POINT(lng lat)"`
            CreatedAt:   time.Now(),
            UpdatedAt:   time.Now(),
        },
    }

    // Insert the dummy data into the wells table
    for _, ship := range ships {
        if err := database.DB.Create(&ship).Error; err != nil {
            log.Fatalf("Failed to seed wells: %v", err)
        }
    }

    log.Println("Wells seeded successfully!")
}
