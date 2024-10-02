package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database"
)

// Get all harbors
func GetHarbors(c echo.Context) error {
    var harbors []models.Harbor
    if result := database.DB.Find(&harbors); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, harbors)
}

// Create a harbor
func CreateHarbor(c echo.Context) error {
    harbor := new(models.Harbor)
    if err := c.Bind(harbor); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if result := database.DB.Create(harbor); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusCreated, harbor)
}

// Get a specific harbor by ID
func GetHarbor(c echo.Context) error {
    id := c.Param("id")
    var harbor models.Harbor
    if result := database.DB.First(&harbor, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, result.Error)
    }
    return c.JSON(http.StatusOK, harbor)
}

// Update a harbor
func UpdateHarbor(c echo.Context) error {
    id := c.Param("id")
    var harbor models.Harbor
    if result := database.DB.First(&harbor, id); result.Error != nil {
        return c.JSON(http.StatusNotFound, result.Error)
    }
    if err := c.Bind(&harbor); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    if result := database.DB.Save(&harbor); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.JSON(http.StatusOK, harbor)
}

// Delete a harbor
func DeleteHarbor(c echo.Context) error {
    id := c.Param("id")
    if result := database.DB.Delete(&models.Harbor{}, id); result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
    return c.NoContent(http.StatusNoContent)
}
