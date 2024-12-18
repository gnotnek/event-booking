package export

import (
	"event-booking/internal/entity"
	"event-booking/internal/export/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// this unit test got nill pointer error, need to work on rabbitmq connection testing

func TestExportAllEvent(t *testing.T) {
	mockEventRepo := mocks.NewEventRepository(t)
	mockBookingRepo := mocks.NewBookingRepository(t)

	mockEvents := []entity.Event{
		{
			ID:            uuid.New(),
			Name:          "Test Event 1",
			Location:      "Test Location 1",
			StartDate:     time.Now(),
			EndDate:       time.Now().Add(time.Hour * 2),
			Price:         100000,
			TotalSeat:     100,
			AvailableSeat: 100,
			Category:      "Test Category 1",
		},
		{
			ID:            uuid.New(),
			Name:          "Test Event 2",
			Location:      "Test Location 2",
			StartDate:     time.Now(),
			EndDate:       time.Now().Add(time.Hour * 2),
			Price:         100000,
			TotalSeat:     100,
			AvailableSeat: 100,
			Category:      "Test Category 2",
		},
	}

	t.Run("export all event successfully", func(t *testing.T) {
		mockEventRepo.On("FindAll").Return(mockEvents, nil).Once()

		svc := NewService(nil, mockEventRepo, mockBookingRepo)
		err := svc.ExportAllEvent()
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, len(mockEvents), len(mockEvents))
	})

	t.Run("export all event failed", func(t *testing.T) {
		mockEventRepo.On("FindAll").Return(nil, assert.AnError).Once()

		svc := NewService(nil, mockEventRepo, mockBookingRepo)
		err := svc.ExportAllEvent()
		if err == nil {
			t.Error("expected error; got nil")
		}

		assert.Equal(t, 0, len(mockEvents))
	})
}

func TestExportAllBookingByUserID(t *testing.T) {
	mockEventRepository := mocks.NewEventRepository(t)
	mockBookingRepository := mocks.NewBookingRepository(t)

	mockBooking := entity.Booking{
		ID:         uuid.New(),
		EventID:    uuid.New(),
		UserID:     uuid.New(),
		Quantity:   1,
		TotalPrice: 100000,
	}

	t.Run("export booking by id successfully", func(t *testing.T) {
		mockBookingRepository.On("FindByUserID", mockBooking.UserID.String()).Return([]entity.Booking{mockBooking}, nil).Once()

		svc := NewService(nil, mockEventRepository, mockBookingRepository)
		err := svc.ExportAllBookingByUser(mockBooking.UserID.String())
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}
	})

	t.Run("export booking by id failed", func(t *testing.T) {
		mockBookingRepository.On("FindByUserID", mockBooking.UserID.String()).Return(nil, assert.AnError).Once()

		svc := NewService(nil, mockEventRepository, mockBookingRepository)
		err := svc.ExportAllBookingByUser(mockBooking.UserID.String())
		if err == nil {
			t.Error("expected error; got nil")
		}
	})
}
