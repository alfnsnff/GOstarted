package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "GOstarted/models"
    "GOstarted/database"
)

// GetTicketClasses retrieves all ticket classes
func GetTicketClasses(c *gin.Context) {
    var ticketClasses []models.TicketClass
    result := database.DB.Find(&ticketClasses)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ticketClasses)
}

// CreateTicketClass adds a new ticket class
func CreateTicketClass(c *gin.Context) {
    var ticketClass models.TicketClass
    if err := c.ShouldBindJSON(&ticketClass); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    result := database.DB.Create(&ticketClass)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusCreated, ticketClass)
}

// GetTicketClassByID retrieves a ticket class by ID
func GetTicketClassByID(c *gin.Context) {
    id := c.Param("id")
    var ticketClass models.TicketClass
    result := database.DB.First(&ticketClass, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ticketClass)
}

// UpdateTicketClass modifies an existing ticket class by ID
func UpdateTicketClass(c *gin.Context) {
    id := c.Param("id")
    var ticketClass models.TicketClass
    result := database.DB.First(&ticketClass, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    if err := c.ShouldBindJSON(&ticketClass); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    result = database.DB.Save(&ticketClass)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ticketClass)
}

// DeleteTicketClass removes a ticket class by ID
func DeleteTicketClass(c *gin.Context) {
    id := c.Param("id")
    result := database.DB.Delete(&models.TicketClass{}, id)
    if result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
