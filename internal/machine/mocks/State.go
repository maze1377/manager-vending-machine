// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	models "github.com/maze1377/manager-vending-machine/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// State is an autogenerated mock type for the State type
type State struct {
	mock.Mock
}

// AddItem provides a mock function with given fields: product
func (_m *State) AddItem(product *models.Product) error {
	ret := _m.Called(product)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Product) error); ok {
		r0 = rf(product)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DispenseProduct provides a mock function with given fields: productName
func (_m *State) DispenseProduct(productName string) error {
	ret := _m.Called(productName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(productName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertMoney provides a mock function with given fields: coin
func (_m *State) InsertMoney(coin float32) error {
	ret := _m.Called(coin)

	var r0 error
	if rf, ok := ret.Get(0).(func(float32) error); ok {
		r0 = rf(coin)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectProduct provides a mock function with given fields: productName
func (_m *State) SelectProduct(productName string) error {
	ret := _m.Called(productName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(productName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewState interface {
	mock.TestingT
	Cleanup(func())
}

// NewState creates a new instance of State. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewState(t mockConstructorTestingTNewState) *State {
	mock := &State{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
