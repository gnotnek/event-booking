package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title       string    `gorm:"not null"`
	Description string
	Location    string
	Date        time.Time `gorm:"not null"`
	Quota       int       `gorm:"not null"`
	Bookings    []Booking `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
