package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	EventID   uuid.UUID `json:"event_id"`
	UserID    uuid.UUID `json:"user_id"`
	Review    string    `json:"review"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time
	UpdatedAt time.Time
	User      User  `json:"user"`
	Event     Event `json:"event"`
}
