package models

import "time"

type Pathway struct {
    ID          uint      `gorm:"primaryKey"`
    ShipID      uint      `json:"ship_id"`        // Ship taking this route
    FromHarborID uint     `json:"from_harbor_id"` // Departure harbor
    ToHarborID  uint      `json:"to_harbor_id"`   // Destination harbor
    TravelTime  string    `json:"travel_time"`    // Estimated travel time (optional)
    Ship        Ship      `gorm:"foreignKey:ShipID"`
    FromHarbor  Harbor    `gorm:"foreignKey:FromHarborID"`
    ToHarbor    Harbor    `gorm:"foreignKey:ToHarborID"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
