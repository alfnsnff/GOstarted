package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "GOstarted/models"
    "GOstarted/database"
)

// Get all harbors
func GetHarbors(c *gin.Context) {
    var harbors []models.Harbor
    if result := database.DB.Find(&harbors); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, harbors)
}

// Create a harbor
func CreateHarbor(c *gin.Context) {
    harbor := new(models.Harbor)
    if err := c.ShouldBindJSON(harbor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Create(harbor); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusCreated, harbor)
}

// Get a specific harbor by ID
func GetHarbor(c *gin.Context) {
    id := c.Param("id")
    var harbor models.Harbor
    if result := database.DB.First(&harbor, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, harbor)
}

// Update a harbor
func UpdateHarbor(c *gin.Context) {
    id := c.Param("id")
    var harbor models.Harbor
    if result := database.DB.First(&harbor, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    if err := c.ShouldBindJSON(&harbor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Save(&harbor); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, harbor)
}

// Delete a harbor
func DeleteHarbor(c *gin.Context) {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Harbor{}, id); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
