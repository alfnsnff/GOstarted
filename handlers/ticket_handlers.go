package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database"
)

// Get all tickets
func GetTickets(c echo.Context) error {
    var tickets []models.Ticket
    if result := database.DB.Preload("Ship").Preload("Pathway").Preload("TicketClass").Find(&tickets); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, tickets)
}

// Create a ticket
func CreateTicket(c echo.Context) error {
    ticket := new(models.Ticket)
    if err := c.Bind(ticket); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if result := database.DB.Create(ticket); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusCreated, ticket)
}

// Get a specific ticket by ID
func GetTicket(c echo.Context) error {
    id := c.Param("id")
    var ticket models.Ticket
    if result := database.DB.Preload("Ship").Preload("Route").Preload("TicketClass").First(&ticket, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, result.Error)
    }
    return c.JSON(http.StatusOK, ticket)
}

// Update a ticket
func UpdateTicket(c echo.Context) error {
    id := c.Param("id")
    var ticket models.Ticket
    if result := database.DB.First(&ticket, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, result.Error)
    }
    if err := c.Bind(&ticket); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error)
    }
    if result := database.DB.Save(&ticket); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, ticket)
}

// Delete a ticket
func DeleteTicket(c echo.Context) error {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Ticket{}, id); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.NoContent(http.StatusNoContent)
}
