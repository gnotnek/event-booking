package entity

import "github.com/google/uuid"

type Booking struct {
	ID        uuid.UUID `json:"id"`
	EventID   uuid.UUID `json:"event_id"`
	UserID    uuid.UUID `json:"user_id"`
	Quantity  int       `json:"quantity"`
	Reference string    `json:"reference"`
	Status    string    `json:"status"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	User      []User    `json:"user"`
	Event     []Event   `json:"event"`
}
