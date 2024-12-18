package event

import (
	"event-booking/internal/entity"
	"fmt"

	"github.com/rs/zerolog/log"
)

//go:generate mockery --case snake --name Repository
type Repository interface {
	Create(event *entity.Event) (*entity.Event, error)
	Save(event *entity.Event) (*entity.Event, error)
	FindAll() ([]entity.Event, error)
	Find(id string) (*entity.Event, error)
	FindByName(name string) (*entity.Event, error)
	FilterByCriteria(criteria map[string]interface{}) ([]entity.Event, error)
	GetBookingsByEventID(eventID string) (entity.Event, error)
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

func (s *Service) CreateEventService(event *entity.Event) (*entity.Event, error) {
	_, err := s.repo.FindByName(event.Name)
	if err == nil {
		log.Error().Msg("event already exists")
		return nil, fmt.Errorf("event already exists")
	}

	event, err = s.repo.Create(event)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return event, nil
}

func (s *Service) SaveEventService(event *entity.Event, newEvent *EventUpdatePayload) (*entity.Event, error) {
	event.Name = newEvent.Name
	event.Location = newEvent.Location
	event.StartDate = newEvent.StartDate
	event.EndDate = newEvent.EndDate
	event.Price = newEvent.Price
	event.TotalSeat = newEvent.TotalSeat
	event.AvailableSeat = newEvent.AvailableSeat

	event, err := s.repo.Save(event)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return event, nil
}

func (s *Service) FindAllEventService() ([]entity.Event, error) {
	events, err := s.repo.FindAll()
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return events, nil
}

func (s *Service) FindEventService(id string) (*entity.Event, error) {
	event, err := s.repo.Find(id)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return event, nil
}

func (s *Service) DeleteEventService(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return nil
}

func (s *Service) FilterEventService(criteria map[string]interface{}) ([]entity.Event, error) {
	events, err := s.repo.FilterByCriteria(criteria)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return events, nil
}

func (s *Service) GetEventBookingsService(eventID string) (entity.Event, error) {
	event, err := s.repo.GetBookingsByEventID(eventID)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return entity.Event{}, err
	}

	return event, nil
}
