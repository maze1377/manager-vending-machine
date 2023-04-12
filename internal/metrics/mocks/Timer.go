// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Timer is an autogenerated mock type for the Timer type
type Timer struct {
	mock.Mock
}

// Done provides a mock function with given fields: started, labels
func (_m *Timer) Done(started time.Time, labels ...string) {
	_va := make([]interface{}, len(labels))
	for _i := range labels {
		_va[_i] = labels[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, started)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewTimer interface {
	mock.TestingT
	Cleanup(func())
}

// NewTimer creates a new instance of Timer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTimer(t mockConstructorTestingTNewTimer) *Timer {
	mock := &Timer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
