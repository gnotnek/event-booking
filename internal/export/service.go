package export

import (
	"context"
	"encoding/json"
	"event-booking/internal/entity"
	"time"

	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog/log"
)

//go:generate mockery --case snake --name EventRepository
type EventRepository interface {
	FindAll() ([]entity.Event, error)
}

//go:generate mockery --case snake --name BookingRepository
type BookingRepository interface {
	FindByUserID(userID string) ([]entity.Booking, error)
}

type Service struct {
	rabbitConn        *amqp091.Connection
	EventRepository   EventRepository
	BookingRepository BookingRepository
}

func NewService(rabbitConn *amqp091.Connection, eventRepo EventRepository, bookingRepo BookingRepository) *Service {
	return &Service{
		rabbitConn:        rabbitConn,
		EventRepository:   eventRepo,
		BookingRepository: bookingRepo,
	}
}

type EventsDataExport struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Location      string    `json:"location"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Price         float64   `json:"price"`
	TotalSeat     int       `json:"total_seat"`
	AvailableSeat int       `json:"available_seat"`
}

func (s *Service) ExportAllEvent() error {
	events, err := s.EventRepository.FindAll()
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	var eventsData []EventsDataExport
	for _, event := range events {
		eventsData = append(eventsData, EventsDataExport{
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

	ch, err := s.rabbitConn.Channel()
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"event_export",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
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

	return nil
}

type BookingsDataExport struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	EventID    uuid.UUID `json:"event_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
}

func (s *Service) ExportAllBookingByUser(userId string) error {
	bookings, err := s.BookingRepository.FindByUserID(userId)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	var bookingsData []BookingsDataExport
	for _, booking := range bookings {
		bookingsData = append(bookingsData, BookingsDataExport{
			ID:         booking.ID,
			UserID:     booking.UserID,
			EventID:    booking.EventID,
			Quantity:   booking.Quantity,
			TotalPrice: booking.TotalPrice,
		})
	}

	ch, err := s.rabbitConn.Channel()
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	queue, err := ch.QueueDeclare(
		"booking_export",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	jsonData, err := json.Marshal(bookingsData)
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

	return nil
}
