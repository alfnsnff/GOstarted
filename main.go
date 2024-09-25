package main

import (
	"net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/database"
	"GOstarted/models"
	"GOstarted/routes"
)

func main() {
    
    database.Connect()
    database.DB.AutoMigrate(&models.Ship{}, &models.Harbor{}, &models.User{}, &models.Ticket{})

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, Echo and PostgreSQL!")
    })
    
	routes.SetupRoutes(e)

    e.Logger.Fatal(e.Start(":8080"))
}
