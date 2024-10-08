package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "GOstarted/models"
    "GOstarted/database"
)

// Get all ships
func GetShips(c *gin.Context) {
    var ships []models.Ship
    if result := database.DB.Find(&ships); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ships)
}

// Create a ship
func CreateShip(c *gin.Context) {
    ship := new(models.Ship)
    if err := c.ShouldBindJSON(ship); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Create(ship); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusCreated, ship)
}

// Get a specific ship by ID
func GetShip(c *gin.Context) {
    id := c.Param("id")
    var ship models.Ship
    if result := database.DB.First(&ship, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ship)
}

// Update a ship
func UpdateShip(c *gin.Context) {
    id := c.Param("id")
    var ship models.Ship
    if result := database.DB.First(&ship, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    if err := c.ShouldBindJSON(&ship); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Save(&ship); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, ship)
}

// Delete a ship
func DeleteShip(c *gin.Context) {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Ship{}, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
