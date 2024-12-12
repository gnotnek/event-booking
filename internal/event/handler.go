package event

import (
	"event-booking/internal/api/responses"
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
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
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
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	eventResponse := responses.EventResponseObject{
		ID:            createdEvent.ID,
		Name:          createdEvent.Name,
		Location:      createdEvent.Location,
		StartDate:     createdEvent.StartDate,
		EndDate:       createdEvent.EndDate,
		Price:         createdEvent.Price,
		TotalSeat:     createdEvent.TotalSeat,
		AvailableSeat: createdEvent.AvailableSeat,
	}

	return c.Status(fiber.StatusCreated).JSON(responses.NewDataResponse(
		"Event created successfully",
		eventResponse,
	))
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
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}

	eventData, err := h.svc.FindEventService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Event not found"))
	}

	newEvent, err := h.svc.SaveEventService(eventData, event)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	updatedEvent := responses.EventResponseObject{
		ID:            newEvent.ID,
		Name:          newEvent.Name,
		Location:      newEvent.Location,
		StartDate:     newEvent.StartDate,
		EndDate:       newEvent.EndDate,
		Price:         newEvent.Price,
		TotalSeat:     newEvent.TotalSeat,
		AvailableSeat: newEvent.AvailableSeat,
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse(
		"Event updated successfully",
		updatedEvent,
	))
}

func (h *httpHandler) FindAllEventHandler(c *fiber.Ctx) error {
	events, err := h.svc.FindAllEventService()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	var eventResponse []responses.EventResponseObject
	for _, event := range events {
		eventResponse = append(eventResponse, responses.EventResponseObject{
			ID:            event.ID,
			Name:          event.Name,
			Location:      event.Location,
			StartDate:     event.StartDate,
			EndDate:       event.EndDate,
			Price:         event.Price,
			TotalSeat:     event.TotalSeat,
			AvailableSeat: event.AvailableSeat,
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse(
		"Events found",
		eventResponse,
	))
}

func (h *httpHandler) FindEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	event, err := h.svc.FindEventService(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Event not found"))
	}

	eventResponse := responses.EventResponseObject{
		ID:            event.ID,
		Name:          event.Name,
		Location:      event.Location,
		StartDate:     event.StartDate,
		EndDate:       event.EndDate,
		Price:         event.Price,
		TotalSeat:     event.TotalSeat,
		AvailableSeat: event.AvailableSeat,
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewDataResponse(
		"Event found",
		eventResponse,
	))
}

func (h *httpHandler) DeleteEventHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(responses.NewErrorResponse("Bad Request"))
	}
	err := h.svc.DeleteEventService(id)
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return c.Status(fiber.StatusNotFound).JSON(responses.NewErrorResponse("Event not found"))
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
		}
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Event deleted successfully"))
}
