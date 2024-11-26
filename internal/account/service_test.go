package account

import (
	"event-booking/internal/account/mocks"
	"event-booking/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUpUser(t *testing.T) {
	mockRepo := mocks.NewRepository(t)
	mockRepo.On("CreateAccount", mock.Anything).Return(nil).Once()

	svc := NewService(mockRepo)
	err := svc.SignUpUserService(&entity.User{
		Email:    "johndoe@gmail.com",
		Password: "password",
	})

	if err != nil {
		t.Errorf("expected error to be nil; got %v", err)
	}

	mockRepo.AssertExpectations(t)
}

func TestSignInUserService(t *testing.T) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	tests := []struct {
		name        string
		email       string
		password    string
		mockReturn  *entity.User
		mockError   error
		expectedErr bool
	}{
		{
			name:     "successful sign in",
			email:    "johndoe@gmail.com",
			password: "password",
			mockReturn: &entity.User{
				Email:    "johndoe@gmail.com",
				Password: string(hashedPassword),
			},
			mockError:   nil,
			expectedErr: false,
		},
		{
			name:        "user not found",
			email:       "janedoe@gmail.com",
			password:    "password",
			mockReturn:  nil,
			mockError:   assert.AnError,
			expectedErr: true,
		},
		{
			name:     "incorrect password",
			email:    "johndoe@gmail.com",
			password: "wrongpassword",
			mockReturn: &entity.User{
				Email:    "johndoe@gmail.com",
				Password: string(hashedPassword),
			},
			mockError:   nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewRepository(t)
			mockRepo.On("FindByEmail", tt.email).Return(tt.mockReturn, tt.mockError).Once()

			svc := NewService(mockRepo)
			_, err := svc.SignInUserService(&entity.User{
				Email:    tt.email,
				Password: tt.password,
			})

			if (err != nil) != tt.expectedErr {
				t.Errorf("expected error to be %v; got %v", tt.expectedErr, err != nil)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
