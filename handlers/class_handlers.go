package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database" // Import your database package
)

// Get all ticket classes
func GetTicketClasses(c echo.Context) error {
    var ticketClasses []models.TicketClass
    if result := database.DB.Find(&ticketClasses); result.Error != nil { // Use database.DB
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, ticketClasses)
}

// Create a ticket class
func CreateTicketClass(c echo.Context) error {
    ticketClass := new(models.TicketClass)
    if err := c.Bind(ticketClass); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if result := database.DB.Create(ticketClass); result.Error != nil { // Use database.DB
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusCreated, ticketClass)
}

// Get a specific ticket class by ID
func GetTicketClass(c echo.Context) error {
    id := c.Param("id")
    var ticketClass models.TicketClass
    if result := database.DB.First(&ticketClass, id); result.Error != nil { // Use database.DB
        return c.JSON(http.StatusNotFound, result.Error)
    }
    return c.JSON(http.StatusOK, ticketClass)
}

// Update a ticket class
func UpdateTicketClass(c echo.Context) error {
    id := c.Param("id")
    var ticketClass models.TicketClass
    if result := database.DB.First(&ticketClass, id); result.Error != nil { // Use database.DB
        return c.JSON(http.StatusNotFound, result.Error)
    }
    if err := c.Bind(&ticketClass); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if result := database.DB.Save(&ticketClass); result.Error != nil { // Use database.DB
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, ticketClass)
}

// Delete a ticket class
func DeleteTicketClass(c echo.Context) error {
    id := c.Param("id")
    if result := database.DB.Delete(&models.TicketClass{}, id); result.Error != nil { // Use database.DB
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.NoContent(http.StatusNoContent)
}
