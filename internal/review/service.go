package review

import (
	"event-booking/internal/entity"

	"github.com/rs/zerolog/log"
)

type Repository interface {
	Create(review *entity.Review) (*entity.Review, error)
	Save(review *entity.Review) (*entity.Review, error)
	FindAll() ([]entity.Review, error)
	Find(id string) (*entity.Review, error)
	FindByEventID(eventID string) ([]entity.Review, error)
	FindByUserID(userID string) ([]entity.Review, error)
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

func (s *Service) CreateReviewService(review *entity.Review) (*entity.Review, error) {
	review, err := s.repo.Create(review)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (s *Service) SaveReviewService(review *entity.Review) (*entity.Review, error) {
	_, err := s.repo.Find(review.ID.String())
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	review, err = s.repo.Save(review)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return review, nil
}

func (s *Service) FindAllReviewService() ([]entity.Review, error) {
	reviews, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *Service) FindReviewService(id string) (*entity.Review, error) {
	review, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (s *Service) FindReviewByEventIDService(eventID string) ([]entity.Review, error) {
	reviews, err := s.repo.FindByEventID(eventID)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *Service) FindReviewByUserIDService(userID string) ([]entity.Review, error) {
	reviews, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}

func (s *Service) DeleteReviewService(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
