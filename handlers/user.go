package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

// GET /users
func GetUsers(c echo.Context) error {
    // Logika untuk mendapatkan daftar user
    return c.JSON(http.StatusOK, "Get all users")
}

// POST /users
func CreateUser(c echo.Context) error {
    // Logika untuk membuat user baru
    return c.JSON(http.StatusCreated, "User created")
}

// GET /users/:id
func GetUserByID(c echo.Context) error {
    id := c.Param("id")  // Mengambil ID dari URL parameter
    // Logika untuk mendapatkan user berdasarkan ID
    return c.JSON(http.StatusOK, "Get user with ID: "+id)
}

// PUT /users/:id
func UpdateUser(c echo.Context) error {
    id := c.Param("id")  // Mengambil ID dari URL parameter
    // Logika untuk mengupdate user berdasarkan ID
    return c.JSON(http.StatusOK, "Update user with ID: "+id)
}

// DELETE /users/:id
func DeleteUser(c echo.Context) error {
    id := c.Param("id")  // Mengambil ID dari URL parameter
    // Logika untuk menghapus user berdasarkan ID
    return c.JSON(http.StatusOK, "Delete user with ID: "+id)
}
