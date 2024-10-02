package models

import "time"

type Harbor struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `json:"name"`     // Name of the harbor
    Location  string    `json:"location"` // Location of the harbor
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
