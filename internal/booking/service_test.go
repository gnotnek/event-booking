package booking

import (
	"event-booking/internal/booking/mocks"
	"event-booking/internal/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookingService(t *testing.T) {
	mockBookingRepo := mocks.NewRepository(t)
	mockEventRepo := mocks.NewEventRepository(t)

	mockRequest := &entity.Booking{
		EventID:  uuid.New(),
		UserID:   uuid.New(),
		Quantity: 2,
	}

	mockEvent := &entity.Event{
		ID:            mockRequest.EventID,
		AvailableSeat: 10,
		Price:         100,
	}

	expectedBooking := &entity.Booking{
		ID:         uuid.New(),
		EventID:    mockRequest.EventID,
		UserID:     mockRequest.UserID,
		Quantity:   mockRequest.Quantity,
		TotalPrice: float64(float64(mockRequest.Quantity) * mockEvent.Price),
	}

	t.Run("create booking successfully", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		mockBookingRepo.On("Create", mockRequest).Return(expectedBooking, nil).Once()

		mockEvent.AvailableSeat -= mockRequest.Quantity

		mockEventRepo.On("Save", mockEvent).Return(mockEvent, nil).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		booking, err := svc.CreateBookingService(mockRequest)
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, expectedBooking, booking)
	})

	t.Run("not enough seat available", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		mockRequest.Quantity = 20

		svc := NewService(mockBookingRepo, mockEventRepo)
		_, err := svc.CreateBookingService(mockRequest)
		assert.Equal(t, "not enough seat available", err.Error())
	})

	t.Run("find event error", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		_, err := svc.CreateBookingService(mockRequest)
		assert.Equal(t, assert.AnError, err)
	})

	t.Run("create booking error", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		mockBookingRepo.On("Create", mockRequest).Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		_, err := svc.CreateBookingService(mockRequest)
		assert.Equal(t, assert.AnError, err)
	})

	t.Run("save event error", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		mockBookingRepo.On("Create", mockRequest).Return(expectedBooking, nil).Once()

		mockEvent.AvailableSeat -= mockRequest.Quantity

		mockEventRepo.On("Save", mockEvent).Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		_, err := svc.CreateBookingService(mockRequest)
		assert.Equal(t, assert.AnError, err)
	})
}

func TestSaveBookingService(t *testing.T) {
	mockBookingRepo := mocks.NewRepository(t)

	mockRequest := &entity.Booking{
		ID:       uuid.New(),
		EventID:  uuid.New(),
		UserID:   uuid.New(),
		Quantity: 3,
	}

	mockRequestUpdate := &BookingInputPayload{
		EventID:  uuid.New(),
		UserID:   uuid.New(),
		Quantity: 2,
	}

	expectedBooking := &entity.Booking{
		ID:         mockRequest.ID,
		EventID:    mockRequest.EventID,
		UserID:     mockRequest.UserID,
		Quantity:   mockRequest.Quantity,
		TotalPrice: 200,
	}

	t.Run("save booking successfully", func(t *testing.T) {
		mockBookingRepo.On("Save", mockRequest).Return(expectedBooking, nil).Once()

		svc := NewService(mockBookingRepo, nil)
		booking, err := svc.SaveBookingService(mockRequest.ID.String(), *mockRequestUpdate)
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, expectedBooking, booking)
	})

	t.Run("save booking error", func(t *testing.T) {
		mockBookingRepo.On("Save", mockRequest).Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, nil)
		_, err := svc.SaveBookingService(mockRequest.ID.String(), *mockRequestUpdate)
		assert.Equal(t, assert.AnError, err)
	})
}

func TestFindAllBookingService(t *testing.T) {
	mockBookingRepo := mocks.NewRepository(t)

	mockBookings := []entity.Booking{
		{
			ID:       uuid.New(),
			EventID:  uuid.New(),
			UserID:   uuid.New(),
			Quantity: 2,
		},
		{
			ID:       uuid.New(),
			EventID:  uuid.New(),
			UserID:   uuid.New(),
			Quantity: 3,
		},
	}

	t.Run("find all booking successfully", func(t *testing.T) {
		mockBookingRepo.On("FindAll").Return(mockBookings, nil).Once()

		svc := NewService(mockBookingRepo, nil)
		bookings, err := svc.FindAllBookingService()
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, mockBookings, bookings)
	})

	t.Run("find all booking error", func(t *testing.T) {
		mockBookingRepo.On("FindAll").Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, nil)
		_, err := svc.FindAllBookingService()
		assert.Equal(t, assert.AnError, err)
	})
}

func TestFindBookingByIDService(t *testing.T) {
	mockBookingRepo := mocks.NewRepository(t)

	mockRequest := &entity.Booking{
		ID:       uuid.New(),
		EventID:  uuid.New(),
		UserID:   uuid.New(),
		Quantity: 2,
	}

	t.Run("booking found", func(t *testing.T) {
		mockBookingRepo.On("Find", mockRequest.ID.String()).Return(mockRequest, nil).Once()

		svc := NewService(mockBookingRepo, nil)
		booking, err := svc.FindBookingService(mockRequest.ID.String())
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, mockRequest, booking)
	})

	t.Run("booking not found", func(t *testing.T) {
		mockBookingRepo.On("Find", mockRequest.ID.String()).Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, nil)
		_, err := svc.FindBookingService(mockRequest.ID.String())
		assert.Equal(t, assert.AnError, err)
	})
}

func TestDeleteBookingService(t *testing.T) {
	mockBookingRepo := mocks.NewRepository(t)
	mockEventRepo := mocks.NewEventRepository(t)

	mockRequest := &entity.Booking{
		ID:       uuid.New(),
		EventID:  uuid.New(),
		UserID:   uuid.New(),
		Quantity: 2,
	}

	mockEvent := &entity.Event{
		ID:            mockRequest.EventID,
		AvailableSeat: 10,
	}

	t.Run("delete booking successfully", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		mockEvent.AvailableSeat += mockRequest.Quantity

		mockBookingRepo.On("Delete", mockRequest.ID.String()).Return(nil).Once()

		mockEventRepo.On("Save", mockEvent).Return(mockEvent, nil).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		err := svc.DeleteBookingService(mockRequest.ID.String(), mockRequest)
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}
	})

	t.Run("find event error", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		err := svc.DeleteBookingService(mockRequest.ID.String(), mockRequest)
		assert.Equal(t, assert.AnError, err)
	})

	t.Run("delete booking error", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		mockEvent.AvailableSeat += mockRequest.Quantity

		mockBookingRepo.On("Delete", mockRequest.ID.String()).Return(assert.AnError).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		err := svc.DeleteBookingService(mockRequest.ID.String(), mockRequest)
		assert.Equal(t, assert.AnError, err)
	})

	t.Run("save event error", func(t *testing.T) {
		mockEventRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		mockEvent.AvailableSeat += mockRequest.Quantity

		mockBookingRepo.On("Delete", mockRequest.ID.String()).Return(nil).Once()

		mockEventRepo.On("Save", mockEvent).Return(nil, assert.AnError).Once()

		svc := NewService(mockBookingRepo, mockEventRepo)
		err := svc.DeleteBookingService(mockRequest.ID.String(), mockRequest)
		assert.Equal(t, assert.AnError, err)
	})
}
