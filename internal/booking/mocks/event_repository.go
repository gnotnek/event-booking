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

// Find provides a mock function with given fields: id
func (_m *EventRepository) Find(id string) (*entity.Event, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 *entity.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Event, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Event); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: event
func (_m *EventRepository) Save(event *entity.Event) (*entity.Event, error) {
	ret := _m.Called(event)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *entity.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Event) (*entity.Event, error)); ok {
		return rf(event)
	}
	if rf, ok := ret.Get(0).(func(*entity.Event) *entity.Event); ok {
		r0 = rf(event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Event) error); ok {
		r1 = rf(event)
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
