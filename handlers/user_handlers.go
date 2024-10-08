package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "GOstarted/models"
    "GOstarted/database"
)

// GetUsers retrieves all users
func GetUsers(c *gin.Context) {
    var users []models.User
    if err := database.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}

// CreateUser adds a new user
func CreateUser(c *gin.Context) {
    user := models.User{}
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Create(&user); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusCreated, user)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// UpdateUser modifies an existing user by ID
func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if result := database.DB.Save(&user); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, user)
}

// DeleteUser removes a user by ID
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if result := database.DB.Delete(&models.User{}, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}
