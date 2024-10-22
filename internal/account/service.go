package account

import (
	"event-booking/internal/entity"

	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	CreateAccount(user *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) SignUpUserService(user *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	return s.repo.CreateAccount(user)
}

func (s *Service) SignInUserService(user *entity.User) (*entity.User, error) {
	userDB, err := s.repo.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password))
	if err != nil {
		return nil, err
	}

	return userDB, nil
}
