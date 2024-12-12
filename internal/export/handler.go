package export

import (
	"event-booking/internal/api/responses"

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

func (h *httpHandler) ExportAllEventHandler(c *fiber.Ctx) error {
	err := h.svc.ExportAllEvent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Export event success"))
}

func (h *httpHandler) ExportBookingHandler(c *fiber.Ctx) error {
	bookingID := c.Params("id")
	err := h.svc.ExportAllBookingByUser(bookingID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Export booking success"))
}
