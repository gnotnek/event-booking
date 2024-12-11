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
