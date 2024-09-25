package models

import (
    "time"
)

type Ticket struct {
    ID        uint    `gorm:"primaryKey"`
    ShipID    uint    `json:"ship_id"`   // Mengaitkan dengan model Ships
    Class     string  `json:"class"`     // Kelas tiket
    Price     float64 `json:"price"`     // Harga tiket
    Passenger string  `json:"passenger"` // Nama penumpang
    Quantity  int     `json:"quantity"`   // Jumlah tiket
	CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
