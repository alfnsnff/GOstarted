package models

import "time"

type Ship struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `json:"name"`     // Name of the ship
    Capacity  int       `json:"capacity"` // Passenger capacity
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
