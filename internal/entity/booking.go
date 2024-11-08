package entity

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;constraint:OnDelete:CASCADE;"`
	EventID   uuid.UUID `gorm:"type:uuid;not null;constraint:OnDelete:CASCADE;"`
	Quantity  int       `gorm:"not null"` // Number of tickets
	Reference string    `gorm:"unique;not null"`
	Status    string
	User      User  `gorm:"foreignKey:UserID"`
	Event     Event `gorm:"foreignKey:EventID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
