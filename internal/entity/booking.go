package entity

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	EventID    uuid.UUID `json:"event_id" gorm:"type:uuid;not null"`
	Quantity   int       `json:"quantity" gorm:"not null"`
	TotalPrice float64   `json:"total_price" gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	User       User  `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Event      Event `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;"`
}
