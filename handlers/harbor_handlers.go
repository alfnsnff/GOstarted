package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "GOstarted/models"
    "GOstarted/database"
)

// GetHarbors retrieves all harbors
func GetHarbors(c echo.Context) error {
    var harbors []models.Harbor
    if err := database.DB.Find(&harbors).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, harbors)
}

// CreateHarbor adds a new harbor
func CreateHarbor(c echo.Context) error {
    harbor := models.Harbor{}
    if err := c.Bind(&harbor); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Create(&harbor)
    return c.JSON(http.StatusCreated, harbor)
}

// UpdateHarbor modifies an existing harbor by ID
func UpdateHarbor(c echo.Context) error {
    id := c.Param("id")
    var harbor models.Harbor
    if err := database.DB.First(&harbor, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    if err := c.Bind(&harbor); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }
    database.DB.Save(&harbor)
    return c.JSON(http.StatusOK, harbor)
}

// DeleteHarbor removes a harbor by ID
func DeleteHarbor(c echo.Context) error {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Harbor{}, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
