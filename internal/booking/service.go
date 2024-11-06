package booking

import (
	"event-booking/internal/entity"

	"github.com/rs/zerolog/log"
)

type Repository interface {
	Create(booking *entity.Booking) (*entity.Booking, error)
	Save(booking *entity.Booking) (*entity.Booking, error)
	Find(id string) (*entity.Booking, error)
	FindAll() ([]entity.Booking, error)
	FindByUserID(userID string) ([]entity.Booking, error)
	FindByEventID(eventID string) ([]entity.Booking, error)
	Delete(id string) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateBookingService(booking *entity.Booking) (*entity.Booking, error) {
	booking, err := s.repo.Create(booking)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return booking, nil
}

func (s *Service) SaveBookingService(booking *entity.Booking) (*entity.Booking, error) {
	booking, err := s.repo.Save(booking)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return booking, nil
}

func (s *Service) FindAllBookingService() ([]entity.Booking, error) {
	bookings, err := s.repo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return bookings, nil
}

func (s *Service) FindByUserIDBookingService(userID string) ([]entity.Booking, error) {
	bookings, err := s.repo.FindByUserID(userID)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return bookings, nil
}

func (s *Service) FindByEventIDBookingService(eventID string) ([]entity.Booking, error) {
	bookings, err := s.repo.FindByEventID(eventID)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return bookings, nil
}

func (s *Service) FindBookingService(id string) (*entity.Booking, error) {
	booking, err := s.repo.Find(id)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return booking, nil
}

func (s *Service) DeleteBookingService(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return nil
}
