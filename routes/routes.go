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
}
