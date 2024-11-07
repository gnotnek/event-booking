package entity

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid;not null;"`
	EventID   uuid.UUID `json:"event_id" gorm:"type:uuid;not null;"`
	Quantity  int       `json:"quantity"` // Number of tickets
	Reference string    `json:"reference"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      *User     `json:"user"`
	Events    []Event   `json:"events" gorm:"many2many:booking_events;"`
}
