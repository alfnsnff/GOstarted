package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "GOstarted/models"
    "GOstarted/database"
)

// Get all tickets
func GetTickets(c *gin.Context) {
    var tickets []models.Ticket
    if result := database.DB.Preload("Ship").Preload("Pathway").Preload("TicketClass").Find(&tickets); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, tickets)
}

// Create a ticket
func CreateTicket(c *gin.Context) {
    ticket := new(models.Ticket)
    if err := c.ShouldBindJSON(ticket); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Create(ticket); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusCreated, ticket)
}

// Get a specific ticket by ID
func GetTicket(c *gin.Context) {
    id := c.Param("id")
    var ticket models.Ticket
    if result := database.DB.Preload("Ship").Preload("Route").Preload("TicketClass").First(&ticket, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ticket)
}

// Update a ticket
func UpdateTicket(c *gin.Context) {
    id := c.Param("id")
    var ticket models.Ticket
    if result := database.DB.First(&ticket, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    if err := c.ShouldBindJSON(&ticket); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Save(&ticket); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ticket)
}

// Delete a ticket
func DeleteTicket(c *gin.Context) {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Ticket{}, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
