package account

import (
	"event-booking/internal/entity"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

//go:generate mockery --case snake --name Repository

type Repository interface {
	CreateAccount(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	UpdateUser(user *entity.User) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SignUpUserService(user *entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	user.Password = string(hashedPassword)

	err = s.repo.CreateAccount(user)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return nil
}

func (s *Service) UpdateUserService(user *entity.User) error {
	userDB, err := s.repo.FindByEmail(user.Email)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	user.ID = userDB.ID

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	user.Password = string(hashedPassword)

	err = s.repo.UpdateUser(user)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return nil
}

func (s *Service) SignInUserService(user *entity.User) (*entity.User, error) {
	userDB, err := s.repo.FindByEmail(user.Email)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return userDB, nil
}

func (s *Service) FindByIDService(id string) (*entity.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return nil, err
	}

	return user, nil
}
