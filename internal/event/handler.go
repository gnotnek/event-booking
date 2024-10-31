package event

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

func (h *httpHandler) CreateEventHandler(c *fiber.Ctx) error {
	event := new(entity.Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	event, err := h.svc.CreateEventService(event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event created successfully",
		"event":   event,
	})
}

func (h *httpHandler) SaveEventHandler(c *fiber.Ctx) error {
	event := new(entity.Event)
	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	event, err := h.svc.SaveEventService(event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event updated successfully",
		"event":   event,
	})
}

func (h *httpHandler) FindAllEventHandler(c *fiber.Ctx) error {
	events, err := h.svc.FindAllEventService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"events": events,
	})
}

func (h *httpHandler) FindEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	event, err := h.svc.FindEventService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Event not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"event": event,
	})
}

func (h *httpHandler) DeleteEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	err := h.svc.DeleteEventService(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event deleted successfully",
	})
}
