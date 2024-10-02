package models

import "time"

type Ticket struct {
    ID              uint        `gorm:"primaryKey"`
    ShipID          uint        `json:"ship_id"`         // Ship ID
    PathwayID       uint        `json:"Pathway_id"`        // Pathway for this ticket
    TicketClassID   uint        `json:"ticket_class_id"` // Reference to TicketClass model
    Passenger       string      `json:"passenger"`       // Name of the passenger
    Ship            Ship        `gorm:"foreignKey:ShipID"`
    Pathway         Pathway     `gorm:"foreignKey:PathwayID"`
    TicketClass     TicketClass `gorm:"foreignKey:TicketClassID"` // Relation to TicketClass
    CreatedAt       time.Time   `json:"created_at"`
    UpdatedAt       time.Time   `json:"updated_at"`
}
