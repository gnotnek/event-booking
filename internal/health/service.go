package health

import (
	"event-booking/internal/entity"

	"github.com/rs/zerolog/log"
)

// go:generate mockery --name Repository
type Repository interface {
	CheckDatabase() error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Check() (*entity.HealthComponent, bool) {
	healthComponent := &entity.HealthComponent{
		Database: entity.HealthStateOK,
	}

	if err := s.repo.CheckDatabase(); err != nil {
		log.Error().Msgf("check database error: %s", err.Error())
		healthComponent.Database = entity.HealthStateFail
	}

	isHealthy := healthComponent.Database == entity.HealthStateOK

	return healthComponent, isHealthy
}
