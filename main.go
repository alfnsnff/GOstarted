package main

import (
    "github.com/labstack/echo/v4"
    "GOstarted/database"
	// "GOstarted/seeds"
	"GOstarted/routes"
)

func main() {
    
    database.Connect()
    // seeds.SeedShips()

    e := echo.New()
	routes.SetupRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}
