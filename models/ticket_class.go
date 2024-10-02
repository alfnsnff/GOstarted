package models

import "time"

// TicketClass represents the class of the ticket (e.g., Economy, Business, First-Class).
type TicketClass struct {
    ID        uint      `gorm:"primaryKey"`
    ClassName string    `json:"class_name"` // Name of the class (e.g., Economy, Business)
    Price     float64   `json:"price"`      // Price for the class
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
