package health

import (
	"context"
	"event-booking/internal/entity"
	"event-booking/internal/health/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name                  string
		mockDatabaseError     error
		expectedDatabaseState entity.HealthState
		expectedRedisState    entity.HealthState
		expectedHealthy       bool
	}{
		{
			name:                  "database unhealthy",
			mockDatabaseError:     assert.AnError,
			expectedDatabaseState: entity.HealthStateFail,
			expectedHealthy:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewRepository(t)
			mockRepo.On("CheckDatabase", mock.Anything).Return(tt.mockDatabaseError).Once()

			svc := &Service{repo: mockRepo}
			healthComponent, isHealthy := svc.Check(context.Background())

			assert.Equal(t, tt.expectedDatabaseState, healthComponent.Database)
			assert.Equal(t, tt.expectedHealthy, isHealthy)

			mockRepo.AssertExpectations(t)
		})
	}
}
