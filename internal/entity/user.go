package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"not null"`
	Role      string    `json:"role" gorm:"not null;default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Bookings  []Booking `gorm:"foreignKey:UserID"`
	Reviews   []Review  `gorm:"foreignKey:UserID"`
}
