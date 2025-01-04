package export

import (
	"context"
	"encoding/json"
	"event-booking/internal/api/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

type httpHandler struct {
	svc        *Service
	rabbitConn *amqp091.Connection
}

func NewHttpHandler(svc *Service, rabittCon *amqp091.Connection) *httpHandler {
	return &httpHandler{
		svc:        svc,
		rabbitConn: rabittCon,
	}
}

func (h *httpHandler) ExportAllEventHandler(c *fiber.Ctx) error {
	eventsData, err := h.svc.ExportAllEvent()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	ch, err := h.rabbitConn.Channel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Error creating rabbitmq channel"))
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"all_export_event",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Error declaring queue"))
	}

	jsonData, err := json.Marshal(eventsData)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	err = ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	})
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Export event success"))
}

func (h *httpHandler) ExportBookingHandler(c *fiber.Ctx) error {
	bookingID := c.Params("id")
	bookings, err := h.svc.ExportAllBookingByUser(bookingID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Internal Server Error"))
	}

	ch, err := h.rabbitConn.Channel()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Error creating rabbitmq channel"))
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"all_export_booking",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Error declaring queue"))
	}

	jsonData, err := json.Marshal(bookings)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
	}

	err = ch.PublishWithContext(context.Background(), "", queue.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	})

	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		c.Status(fiber.StatusInternalServerError).JSON(responses.NewErrorResponse("Error publishing message"))
	}

	return c.Status(fiber.StatusOK).JSON(responses.NewSuccessResponse("Export booking success"))
}
