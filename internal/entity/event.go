package entity

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name          string    `json:"name" gorm:"unique;not null"`
	Location      string    `json:"location"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Price         float64   `json:"price"`
	TotalSeat     int       `json:"total_seat"`
	AvailableSeat int       `json:"available_seat"`
	Category      string    `json:"category"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Bookings      []Booking `gorm:"foreignKey:EventID"`
	Reviews       []Review  `gorm:"foreignKey:EventID"`
}
