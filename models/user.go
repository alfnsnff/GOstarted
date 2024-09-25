package models

import (
    "time"
)

type User struct {
    ID           int       `json:"id"`
    Name         string    `json:"name"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"password_hash"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

// CreateUser saves a new user to the database
func CreateUser(user *User) error {
    // Logic to insert user into the database
    // Example with db.Exec or db.Query using PostgreSQL driver
    return nil
}
