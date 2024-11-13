package booking

import (
	"event-booking/internal/entity"

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
		"booking": newBook,
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
	return nil
}

func (h *httpHandler) CancelBookedEventHandler(c *fiber.Ctx) error {
	return nil
}

func (h *httpHandler) UpdateBookedEventHandler(c *fiber.Ctx) error {
	return nil
}
