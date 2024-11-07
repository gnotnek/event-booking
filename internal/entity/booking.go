package entity

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id"`
	Quantity  int       `json:"quantity"` // Number of tickets
	Reference string    `json:"reference"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user"`
	Events    []Event   `json:"events" gorm:"many2many:booking_events;"`
}
