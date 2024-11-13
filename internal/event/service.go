package event

import (
	"event-booking/internal/entity"

	"github.com/rs/zerolog/log"
)

// go generate mockery -name Repository
type Repository interface {
	Create(event *entity.Event) (*entity.Event, error)
	Save(event *entity.Event) (*entity.Event, error)
	FindAll() ([]entity.Event, error)
	Find(id string) (*entity.Event, error)
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
	event, err := s.repo.Create(event)
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
