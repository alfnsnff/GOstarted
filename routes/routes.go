package routes

import (
    "GOstarted/handlers"
    "github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
    // User endpoints
    e.GET("/users", handlers.GetUsers)
    e.POST("/users", handlers.CreateUser)
    e.GET("/users/:id", handlers.GetUserByID)
    e.PUT("/users/:id", handlers.UpdateUser)
    e.DELETE("/users/:id", handlers.DeleteUser)
    
    e.GET("/ships", handlers.GetShips)
    e.POST("/ships", handlers.CreateShip)
    e.PUT("/ships/:id", handlers.UpdateShip)
    e.DELETE("/ships/:id", handlers.DeleteShip)

    e.GET("/harbors", handlers.GetHarbors)
    e.POST("/harbors", handlers.CreateHarbor)
    e.PUT("/harbors/:id", handlers.UpdateHarbor)
    e.DELETE("/harbors/:id", handlers.DeleteHarbor)

    // Routes for Tickets
    e.GET("/tickets", handlers.GetTickets)
    e.POST("/tickets", handlers.CreateTicket)
    e.PUT("/tickets/:id", handlers.UpdateTicket)
    e.DELETE("/tickets/:id", handlers.DeleteTicket)
}
