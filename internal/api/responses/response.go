package responses

import (
	"time"

	"github.com/google/uuid"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type DataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserResponseObject struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
	Role  string    `json:"role"`
}

type BookingResponseObject struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	EventID    uuid.UUID `json:"event_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type EventResponseObject struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Price         float64   `json:"price"`
	TotalSeat     int       `json:"total_seat"`
	AvailableSeat int       `json:"available_seat"`
	Category      string    `json:"category"`
}

type ReviewResponseObject struct {
	ID        uuid.UUID `json:"id"`
	EventID   int       `json:"event_id"`
	UserID    int       `json:"user_id"`
	Review    string    `json:"review"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewSuccessResponse(message string) *SuccessResponse {
	return &SuccessResponse{
		Message: message,
	}
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}

func NewDataResponse(message string, data interface{}) *DataResponse {
	return &DataResponse{
		Message: message,
		Data:    data,
	}
}
