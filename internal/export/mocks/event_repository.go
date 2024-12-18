// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	entity "event-booking/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// EventRepository is an autogenerated mock type for the EventRepository type
type EventRepository struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *EventRepository) FindAll() ([]entity.Event, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []entity.Event
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]entity.Event, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []entity.Event); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Event)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEventRepository creates a new instance of EventRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventRepository {
	mock := &EventRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
