package event

import (
	"event-booking/internal/entity"
	"time"

	"github.com/gofiber/fiber/v2"
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

type EventInputPayload struct {
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Price         float64   `json:"price"`
	TotalSeat     int       `json:"total_seat"`
	AvailableSeat int       `json:"available_seat"`
}

func (h *httpHandler) CreateEventHandler(c *fiber.Ctx) error {
	event := new(EventInputPayload)
	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	newEvent := &entity.Event{
		Name:          event.Name,
		Location:      event.Location,
		StartDate:     event.StartDate,
		EndDate:       event.EndDate,
		Price:         event.Price,
		TotalSeat:     event.TotalSeat,
		AvailableSeat: event.AvailableSeat,
	}

	createdEvent, err := h.svc.CreateEventService(newEvent)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event created successfully",
		"event":   createdEvent,
	})
}

type EventResponsePayload struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Price         float64   `json:"price"`
	TotalSeat     int       `json:"total_seat"`
	AvailableSeat int       `json:"available_seat"`
}

type EventUpdatePayload struct {
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Price         float64   `json:"price"`
	TotalSeat     int       `json:"total_seat"`
	AvailableSeat int       `json:"available_seat"`
}

func (h *httpHandler) SaveEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	event := new(EventUpdatePayload)
	if err := c.BodyParser(event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}

	eventData, err := h.svc.FindEventService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Event not found",
		})
	}

	newEvent, err := h.svc.SaveEventService(eventData, event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event updated successfully",
		"event": EventResponsePayload{
			ID:            newEvent.ID.String(),
			Name:          newEvent.Name,
			Location:      newEvent.Location,
			StartDate:     newEvent.StartDate,
			EndDate:       newEvent.EndDate,
			Price:         newEvent.Price,
			TotalSeat:     newEvent.TotalSeat,
			AvailableSeat: newEvent.AvailableSeat,
		},
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
		"event": EventResponsePayload{
			ID:            event.ID.String(),
			Name:          event.Name,
			Location:      event.Location,
			StartDate:     event.StartDate,
			EndDate:       event.EndDate,
			Price:         event.Price,
			TotalSeat:     event.TotalSeat,
			AvailableSeat: event.AvailableSeat,
		},
	})
}

func (h *httpHandler) DeleteEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Bad Request",
		})
	}
	err := h.svc.DeleteEventService(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status":  fiber.StatusNotFound,
				"message": "Event not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status":  fiber.StatusInternalServerError,
				"message": "Internal Server Error",
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event deleted successfully",
	})
}
