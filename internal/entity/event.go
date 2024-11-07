package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Quota       int       `json:"quota"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Bookings    []Booking `json:"bookings" gorm:"many2many:booking_events;"`
}
