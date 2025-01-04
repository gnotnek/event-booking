package export

import (
	"event-booking/internal/entity"
	"time"

	"github.com/google/uuid"
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
	EventRepository   EventRepository
	BookingRepository BookingRepository
}

func NewService(eventRepo EventRepository, bookingRepo BookingRepository) *Service {
	return &Service{
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

func (s *Service) ExportAllEvent() ([]EventsDataExport, error) {
	events, err := s.EventRepository.FindAll()
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return []EventsDataExport{}, err
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

	return eventsData, nil
}

type BookingsDataExport struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	EventID    uuid.UUID `json:"event_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
}

func (s *Service) ExportAllBookingByUser(userId string) ([]BookingsDataExport, error) {
	bookings, err := s.BookingRepository.FindByUserID(userId)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return []BookingsDataExport{}, err
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

	return bookingsData, nil
}
