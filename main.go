package main

import (
    "github.com/gin-gonic/gin"
    "GOstarted/database"
	"GOstarted/seeds"
	"GOstarted/routes"
)

func main() {

    // Connect to the database
    database.Connect()

    // Seed ships data
    seeds.SeedShips()

    // Create a new Gin router
    r := gin.Default()

    // Setup routes (adjust routes.SetupRoutes to work with Gin)
    routes.SetupRoutes(r)

    // Start the server on port 8080
    r.Run(":8080")
}
