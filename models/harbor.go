package models

import (
    "time"
)

type Harbor struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
