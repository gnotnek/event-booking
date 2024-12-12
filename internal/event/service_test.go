package event

import (
	"event-booking/internal/entity"
	"event-booking/internal/event/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	mockEvent := &entity.Event{
		ID:            uuid.New(),
		Name:          "Test Event",
		Location:      "Test Location",
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(time.Hour * 2),
		Price:         100000,
		TotalSeat:     100,
		AvailableSeat: 100,
	}

	t.Run("create event successfully", func(t *testing.T) {
		mockRepo.On("Create", mockEvent).Return(mockEvent, nil).Once()
		mockRepo.On("FindByName", mockEvent.Name).Return(nil, nil).Once()

		svc := NewService(mockRepo)
		event, err := svc.CreateEventService(mockEvent)
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, mockEvent, event)
	})

	t.Run("event already exists", func(t *testing.T) {
		mockRepo.On("FindByName", mockEvent.Name).Return(mockEvent, nil).Once()

		svc := NewService(mockRepo)
		_, err := svc.CreateEventService(mockEvent)
		if err == nil {
			t.Error("expected error; got nil")
		}
	})

	t.Run("create event failed", func(t *testing.T) {
		mockRepo.On("Create", mockEvent).Return(nil, assert.AnError).Once()

		svc := NewService(mockRepo)
		_, err := svc.CreateEventService(mockEvent)
		if err == nil {
			t.Error("expected error; got nil")
		}
	})
}

func TestSaveEvent(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	mockEvent := &entity.Event{
		ID:            uuid.New(),
		Name:          "Test Event",
		Location:      "Test Location",
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(time.Hour * 2),
		Price:         100000,
		TotalSeat:     100,
		AvailableSeat: 100,
	}

	newEvent := &EventUpdatePayload{
		Name:          "New Event",
		Location:      "New Location",
		StartDate:     time.Now().Add(time.Hour * 3),
		EndDate:       time.Now().Add(time.Hour * 5),
		Price:         200000,
		TotalSeat:     200,
		AvailableSeat: 200,
	}

	t.Run("save event successfully", func(t *testing.T) {
		mockRepo.On("Save", mockEvent).Return(mockEvent, nil).Once()

		svc := NewService(mockRepo)
		event, err := svc.SaveEventService(mockEvent, newEvent)
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, mockEvent, event)
	})

	t.Run("save event failed", func(t *testing.T) {
		mockRepo.On("Save", mockEvent).Return(nil, assert.AnError).Once()

		svc := NewService(mockRepo)
		_, err := svc.SaveEventService(mockEvent, newEvent)
		if err == nil {
			t.Error("expected error; got nil")
		}
	})
}

func TestFindAllEvent(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	mockEvents := []entity.Event{
		{
			ID:            uuid.New(),
			Name:          "Test Event",
			Location:      "Test Location",
			StartDate:     time.Now(),
			EndDate:       time.Now().Add(time.Hour * 2),
			Price:         100000,
			TotalSeat:     100,
			AvailableSeat: 100,
		},
		{
			ID:            uuid.New(),
			Name:          "Test Event 2",
			Location:      "Test Location 2",
			StartDate:     time.Now().Add(time.Hour * 3),
			EndDate:       time.Now().Add(time.Hour * 5),
			Price:         200000,
			TotalSeat:     200,
			AvailableSeat: 200,
		},
	}

	t.Run("find all event successfully", func(t *testing.T) {
		mockRepo.On("FindAll").Return(mockEvents, nil).Once()

		svc := NewService(mockRepo)
		events, err := svc.FindAllEventService()
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, mockEvents, events)
	})

	t.Run("find all event failed", func(t *testing.T) {
		mockRepo.On("FindAll").Return(nil, assert.AnError).Once()

		svc := NewService(mockRepo)
		_, err := svc.FindAllEventService()
		if err == nil {
			t.Error("expected error; got nil")
		}

		assert.Equal(t, assert.AnError, err)
	})
}

func TestFindEvent(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	mockEvent := &entity.Event{
		ID:            uuid.New(),
		Name:          "Test Event",
		Location:      "Test Location",
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(time.Hour * 2),
		Price:         100000,
		TotalSeat:     100,
		AvailableSeat: 100,
	}

	t.Run("find event successfully", func(t *testing.T) {
		mockRepo.On("Find", mockEvent.ID.String()).Return(mockEvent, nil).Once()

		svc := NewService(mockRepo)
		event, err := svc.FindEventService(mockEvent.ID.String())
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}

		assert.Equal(t, mockEvent, event)
	})

	t.Run("find event failed", func(t *testing.T) {
		mockRepo.On("Find", mockEvent.ID.String()).Return(nil, assert.AnError).Once()

		svc := NewService(mockRepo)
		_, err := svc.FindEventService(mockEvent.ID.String())
		if err == nil {
			t.Error("expected error; got nil")
		}

		assert.Equal(t, assert.AnError, err)
	})
}

func TestDeleteEvent(t *testing.T) {
	mockRepo := mocks.NewRepository(t)

	mockEvent := &entity.Event{
		ID:            uuid.New(),
		Name:          "Test Event",
		Location:      "Test Location",
		StartDate:     time.Now(),
		EndDate:       time.Now().Add(time.Hour * 2),
		Price:         100000,
		TotalSeat:     100,
		AvailableSeat: 100,
	}

	t.Run("delete event successfully", func(t *testing.T) {
		mockRepo.On("Delete", mockEvent.ID.String()).Return(nil).Once()

		svc := NewService(mockRepo)
		err := svc.DeleteEventService(mockEvent.ID.String())
		if err != nil {
			t.Errorf("expected error to be nil; got %v", err)
		}
	})

	t.Run("delete event failed", func(t *testing.T) {
		mockRepo.On("Delete", mockEvent.ID.String()).Return(assert.AnError).Once()

		svc := NewService(mockRepo)
		err := svc.DeleteEventService(mockEvent.ID.String())
		if err == nil {
			t.Error("expected error; got nil")
		}

		assert.Equal(t, assert.AnError, err)
	})
}
