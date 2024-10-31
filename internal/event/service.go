package event

import "event-booking/internal/entity"

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
	return s.repo.Create(event)
}

func (s *Service) SaveEventService(event *entity.Event) (*entity.Event, error) {
	return s.repo.Save(event)
}

func (s *Service) FindAllEventService() ([]entity.Event, error) {
	return s.repo.FindAll()
}

func (s *Service) FindEventService(id string) (*entity.Event, error) {
	return s.repo.Find(id)
}

func (s *Service) DeleteEventService(id string) error {
	return s.repo.Delete(id)
}
