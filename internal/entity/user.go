package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName string
	LastName  string
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Bookings  []Booking `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
