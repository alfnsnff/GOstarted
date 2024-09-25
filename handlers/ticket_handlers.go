package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database"
)

// GetTickets retrieves all tickets
func GetTickets(c echo.Context) error {
    var tickets []models.Ticket
    if err := database.DB.Preload("Ship").Find(&tickets).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, tickets)
}

// CreateTicket adds a new ticket
func CreateTicket(c echo.Context) error {
    ticket := models.Ticket{}
    if err := c.Bind(&ticket); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Create(&ticket)
    return c.JSON(http.StatusCreated, ticket)
}

// UpdateTicket modifies an existing ticket by ID
func UpdateTicket(c echo.Context) error {
    id := c.Param("id")
    var ticket models.Ticket
    if err := database.DB.First(&ticket, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    if err := c.Bind(&ticket); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Save(&ticket)
    return c.JSON(http.StatusOK, ticket)
}

// DeleteTicket removes a ticket by ID
func DeleteTicket(c echo.Context) error {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Ticket{}, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
