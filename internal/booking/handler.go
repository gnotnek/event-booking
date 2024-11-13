package booking

import (
	"event-booking/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type httpHandler struct {
	svc *Service
}

func NewHttpHandler(svc *Service) *httpHandler {
	return &httpHandler{
		svc: svc,
	}
}

type BookingInputPayload struct {
	UserID   uuid.UUID `json:"user_id"`
	EventID  uuid.UUID `json:"event_id"`
	Quantity int       `json:"quantity"`
}

type BookingResponse struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	EventID    uuid.UUID `json:"event_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (h *httpHandler) BookEventHandler(c *fiber.Ctx) error {
	book := new(BookingInputPayload)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	newBook := &entity.Booking{
		UserID:   book.UserID,
		EventID:  book.EventID,
		Quantity: book.Quantity,
	}

	newBook, err := h.svc.CreateBookingService(newBook)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event booked successfully",
		"booking": BookingResponse{
			ID:         newBook.ID,
			UserID:     newBook.UserID,
			EventID:    newBook.EventID,
			Quantity:   newBook.Quantity,
			TotalPrice: newBook.TotalPrice,
			CreatedAt:  newBook.CreatedAt,
			UpdatedAt:  newBook.UpdatedAt,
		},
	})
}

func (h *httpHandler) GetBookedEventsHandler(c *fiber.Ctx) error {
	bookings, err := h.svc.FindAllBookingService()
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  fiber.StatusNotFound,
				"message": "No booking found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bookings": bookings,
	})
}

func (h *httpHandler) GetBookedEventByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := h.svc.FindBookingService(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  fiber.StatusNotFound,
				"message": "Booking not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"booking": book,
	})
}

func (h *httpHandler) CancelBookedEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	book, err := h.svc.FindBookingService(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  fiber.StatusNotFound,
				"message": "Booking not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
	}

	err = h.svc.DeleteBookingService(id, book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Booking canceled successfully",
	})
}

func (h *httpHandler) UpdateBookedEventHandler(c *fiber.Ctx) error {
	return nil
}
