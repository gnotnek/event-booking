package booking

import (
	"event-booking/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	svc *Service
}

func NewHttpHandler(svc *Service) *httpHandler {
	return &httpHandler{
		svc: svc,
	}
}

func (h *httpHandler) BookEventHandler(c *fiber.Ctx) error {
	book := new(entity.Booking)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	book, err := h.svc.CreateBookingService(book)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event booked successfully",
		"booking": book,
	})
}

func (h *httpHandler) GetBookedEventsHandler(c *fiber.Ctx) error {
	return nil
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
