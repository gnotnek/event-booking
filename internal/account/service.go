package account

import (
	"event-booking/internal/email"
	"event-booking/internal/entity"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

//go:generate mockery --case snake --name Repository

type Repository interface {
	CreateAccount(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id string) (*entity.User, error)
	SaveUser(user *entity.User) error
}

type Service struct {
	repo         Repository
	emailService *email.EmailService
}

func NewService(repo Repository, emailService *email.EmailService) *Service {
	return &Service{
		repo:         repo,
		emailService: emailService,
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

	err = s.repo.SaveUser(user)
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

func (s *Service) GenerateVerificationCode(email string) error {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	user.EmailVerificationCode = code
	user.VerificationExpiry = time.Now().Add(1 * time.Hour)
	user.VerificationAttemptsLeft = 3

	if err := s.repo.SaveUser(user); err != nil {
		return fmt.Errorf("failed to save verification code: %v", err)
	}

	err = s.emailService.SendVerificationEmail(user.Email, code)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return nil
}

func (s *Service) ValidateVerificationCode(email, code string) error {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	if time.Now().After(user.VerificationExpiry) {
		return fmt.Errorf("verification code has expired")
	}

	if user.VerificationAttemptsLeft == 0 {
		return fmt.Errorf("no attempts left")
	}

	if user.EmailVerificationCode != code {
		user.VerificationAttemptsLeft--
		if err := s.repo.SaveUser(user); err != nil {
			return fmt.Errorf("failed to save verification attempts: %v", err)
		}
	}

	user.EmailVerificationCode = ""
	user.VerificationExpiry = time.Time{}
	user.VerificationAttemptsLeft = 0
	err = s.repo.SaveUser(user)
	if err != nil {
		log.Error().Err(err).Msg(err.Error())
		return err
	}

	return nil
}
