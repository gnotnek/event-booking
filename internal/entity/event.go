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
	Kouta       int       `json:"available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Booking     []Booking `json:"booking"`
}
