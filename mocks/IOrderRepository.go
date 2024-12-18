// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	dtos "velocity-technical-test/internal/domain/orders/dtos"

	mock "github.com/stretchr/testify/mock"
)

// IOrderRepository is an autogenerated mock type for the IOrderRepository type
type IOrderRepository struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: order
func (_m *IOrderRepository) CreateOrder(order dtos.OrderDTO) (uint, error) {
	ret := _m.Called(order)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(dtos.OrderDTO) (uint, error)); ok {
		return rf(order)
	}
	if rf, ok := ret.Get(0).(func(dtos.OrderDTO) uint); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(dtos.OrderDTO) error); ok {
		r1 = rf(order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOrderItems provides a mock function with given fields: orderItems
func (_m *IOrderRepository) CreateOrderItems(orderItems []dtos.OrderItemDTO) error {
	ret := _m.Called(orderItems)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrderItems")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]dtos.OrderItemDTO) error); ok {
		r0 = rf(orderItems)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOrderWithItems provides a mock function with given fields: orderID
func (_m *IOrderRepository) GetOrderWithItems(orderID uint) (*dtos.OrderDTO, error) {
	ret := _m.Called(orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderWithItems")
	}

	var r0 *dtos.OrderDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*dtos.OrderDTO, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(uint) *dtos.OrderDTO); ok {
		r0 = rf(orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dtos.OrderDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIOrderRepository creates a new instance of IOrderRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIOrderRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IOrderRepository {
	mock := &IOrderRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
