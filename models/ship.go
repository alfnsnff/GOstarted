package models

import (
    "time"
)

type Ship struct {
    ID       uint    `gorm:"primaryKey"`
    Name     string  `json:"name"`
    Capacity int     `json:"capacity"`
    Price    float64 `json:"price"`
    Class    string  `json:"class"`
    HarborID uint    `json:"harbor_id"` // Mengaitkan dengan model Harbor
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
