package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database"
)

// GetUsers retrieves all users
func GetUsers(c echo.Context) error {
    var users []models.User
    if err := database.DB.Find(&users).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, users)
}

// CreateUser adds a new user
func CreateUser(c echo.Context) error {
    user := models.User{}
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Create(&user)
    return c.JSON(http.StatusCreated, user)
}

// GetUserByID retrieves a user by ID
func GetUserByID(c echo.Context) error {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    return c.JSON(http.StatusOK, user)
}

// UpdateUser modifies an existing user by ID
func UpdateUser(c echo.Context) error {
    id := c.Param("id")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Save(&user)
    return c.JSON(http.StatusOK, user)
}

// DeleteUser removes a user by ID
func DeleteUser(c echo.Context) error {
    id := c.Param("id")
    if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
