// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	dtos "velocity-technical-test/internal/domain/products/dtos"

	mock "github.com/stretchr/testify/mock"
)

// IProductUsecase is an autogenerated mock type for the IProductUsecase type
type IProductUsecase struct {
	mock.Mock
}

// GetProducts provides a mock function with no fields
func (_m *IProductUsecase) GetProducts() ([]dtos.ProductDTO, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProducts")
	}

	var r0 []dtos.ProductDTO
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]dtos.ProductDTO, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []dtos.ProductDTO); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dtos.ProductDTO)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProductStock provides a mock function with given fields: productID, newStock
func (_m *IProductUsecase) UpdateProductStock(productID uint, newStock uint) error {
	ret := _m.Called(productID, newStock)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProductStock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(productID, newStock)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIProductUsecase creates a new instance of IProductUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProductUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProductUsecase {
	mock := &IProductUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
