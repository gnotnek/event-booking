package account

import (
	"event-booking/internal/account/mocks"
	"event-booking/internal/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestSignUpUser(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	mockUser := &entity.User{
		Email:    "johndoe@gmail.com",
		Password: "password",
	}

	t.Run("sign up user successfully", func(t *testing.T) {
		mockRepo.On("CreateAccount", mockUser).Return(nil).Once()

		svc := NewService(mockRepo)
		err := svc.SignUpUserService(mockUser)
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		mockRepo.AssertExpectations(t)
	})

	t.Run("sign up user error", func(t *testing.T) {
		mockRepo.On("CreateAccount", mockUser).Return(assert.AnError).Once()

		svc := NewService(mockRepo)
		err := svc.SignUpUserService(mockUser)
		assert.Equal(t, assert.AnError, err)

		mockRepo.AssertExpectations(t)
	})
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

func TestUpdateUserService(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	mockUser := &entity.User{
		ID:       uuid.New(),
		Email:    "johndoe@gmail.com",
		Password: "password",
	}

	mockNewUser := &entity.User{
		ID:       mockUser.ID,
		Email:    "johndoe@gmail.com",
		Password: "newpassword",
	}

	t.Run("update user successfully", func(t *testing.T) {
		mockRepo.On("FindByEmail", mockUser.Email).Return(mockUser, nil).Once()
		mockRepo.On("UpdateUser", mockNewUser).Return(nil).Once()

		svc := NewService(mockRepo)
		err := svc.UpdateUserService(mockNewUser)
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		mockRepo.AssertExpectations(t)
	})

	t.Run("update user error", func(t *testing.T) {
		mockRepo.On("FindByEmail", mockUser.Email).Return(mockUser, nil).Once()
		mockRepo.On("UpdateUser", mockNewUser).Return(assert.AnError).Once()

		svc := NewService(mockRepo)
		err := svc.UpdateUserService(mockNewUser)
		assert.Equal(t, assert.AnError, err)

		mockRepo.AssertExpectations(t)
	})
}
