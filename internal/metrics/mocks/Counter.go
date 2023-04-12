// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Counter is an autogenerated mock type for the Counter type
type Counter struct {
	mock.Mock
}

// Inc provides a mock function with given fields: labels
func (_m *Counter) Inc(labels ...string) {
	_va := make([]interface{}, len(labels))
	for _i := range labels {
		_va[_i] = labels[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

type mockConstructorTestingTNewCounter interface {
	mock.TestingT
	Cleanup(func())
}

// NewCounter creates a new instance of Counter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCounter(t mockConstructorTestingTNewCounter) *Counter {
	mock := &Counter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
