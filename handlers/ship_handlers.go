package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database"
)

// GetShips retrieves all ships
func GetShips(c echo.Context) error {
    var ships []models.Ship
    if err := database.DB.Find(&ships).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, ships)
}

// CreateShip adds a new ship
func CreateShip(c echo.Context) error {
    ship := models.Ship{}
    if err := c.Bind(&ship); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Create(&ship)
    return c.JSON(http.StatusCreated, ship)
}

// UpdateShip modifies an existing ship by ID
func UpdateShip(c echo.Context) error {
    id := c.Param("id")
    var ship models.Ship
    if err := database.DB.First(&ship, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    if err := c.Bind(&ship); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Save(&ship)
    return c.JSON(http.StatusOK, ship)
}

// DeleteShip removes a ship by ID
func DeleteShip(c echo.Context) error {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Ship{}, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
