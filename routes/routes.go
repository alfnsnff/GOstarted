package routes

import (
    "net/http"
    "GOstarted/handlers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Gin and PostgreSQL!")
    })

    // User endpoints
    r.GET("/users", handlers.GetUsers)
    r.POST("/users", handlers.CreateUser)
    r.GET("/users/:id", handlers.GetUserByID)
    r.PUT("/users/:id", handlers.UpdateUser)
    r.DELETE("/users/:id", handlers.DeleteUser)

    // Ship endpoints
    r.GET("/ships", handlers.GetShips)
    r.POST("/ships", handlers.CreateShip)
    r.PUT("/ships/:id", handlers.UpdateShip)
    r.DELETE("/ships/:id", handlers.DeleteShip)

    // Harbor endpoints
    r.GET("/harbors", handlers.GetHarbors)
    r.POST("/harbors", handlers.CreateHarbor)
    r.PUT("/harbors/:id", handlers.UpdateHarbor)
    r.DELETE("/harbors/:id", handlers.DeleteHarbor)

    // Ticket endpoints
    r.GET("/tickets", handlers.GetTickets)
    r.POST("/tickets", handlers.CreateTicket)
    r.PUT("/tickets/:id", handlers.UpdateTicket)
    r.DELETE("/tickets/:id", handlers.DeleteTicket)

    // Ticket Class endpoints
    r.GET("/ticketclasses", handlers.GetTicketClasses)
    r.POST("/ticketclasses", handlers.CreateTicketClass)
    r.PUT("/ticketclasses/:id", handlers.UpdateTicketClass)
    r.DELETE("/ticketclasses/:id", handlers.DeleteTicketClass)
}
