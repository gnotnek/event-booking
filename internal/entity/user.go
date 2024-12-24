package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID                       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name                     string    `json:"name" gorm:"not null"`
	Email                    string    `json:"email" gorm:"unique;not null"`
	Password                 string    `json:"password" gorm:"not null"`
	Role                     string    `json:"role" gorm:"not null;default:'user'"`
	EmailVerificationCode    string    `json:"email_verification_code"`
	VerificationExpiry       time.Time `json:"verification_expiry"`
	VerificationAttemptsLeft int       `json:"verification_attempts_left" gorm:"default:3"`
	CreatedAt                time.Time
	UpdatedAt                time.Time
	Bookings                 []Booking `gorm:"foreignKey:UserID"`
	Reviews                  []Review  `gorm:"foreignKey:UserID"`
}
