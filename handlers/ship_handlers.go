package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database"
)

// Get all ships
func GetShips(c echo.Context) error {
    var ships []models.Ship
    if result := database.DB.Find(&ships); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, ships)
}

// Create a ship
func CreateShip(c echo.Context) error {
    ship := new(models.Ship)
    if err := c.Bind(ship); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if result := database.DB.Create(ship); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusCreated, ship)
}

// Get a specific ship by ID
func GetShip(c echo.Context) error {
    id := c.Param("id")
    var ship models.Ship
    if result := database.DB.First(&ship, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, result.Error)
    }
    return c.JSON(http.StatusOK, ship)
}

// Update a ship
func UpdateShip(c echo.Context) error {
    id := c.Param("id")
    var ship models.Ship
    if result := database.DB.First(&ship, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, result.Error)
    }
    if err := c.Bind(&ship); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error)
    }
    if result := database.DB.Save(&ship); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, ship)
}

// Delete a ship
func DeleteShip(c echo.Context) error {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Ship{}, id); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.NoContent(http.StatusNoContent)
}
