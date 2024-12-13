package usecase

import (
	"testing"
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder(t *testing.T) {
	mockOrderRepo := new(mocks.IOrderRepository)
	mockProductRepo := new(mocks.IProductRepository)
	mockRedisOrder := new(mocks.IRedisOrder)
	orderUseCase := NewOrder(mockOrderRepo, mockProductRepo, mockRedisOrder)

	orderDTO := dtos.OrderDTO{
		Items: []dtos.OrderItemDTO{
			{ProductID: 1, Quantity: 2},
		},
	}

	t.Run("successful order creation", func(t *testing.T) {
		mockRedisOrder.On("KeyExists", mock.Anything).Return(false, nil)
		mockRedisOrder.On("SetJSON", mock.Anything, "IN_PROGRESS", mock.Anything).Return(nil)
		mockProductRepo.On("ProductExists", uint(1)).Return(true, nil)
		mockProductRepo.On("CountProductStock", uint(1)).Return(uint(10), nil)
		mockProductRepo.On("GetProductPrice", uint(1)).Return(100.0, nil)
		mockProductRepo.On("UpdateProductStock", uint(1), uint(8)).Return(nil)
		mockOrderRepo.On("CreateOrder", mock.Anything).Return(uint(1), nil)
		mockOrderRepo.On("CreateOrderItems", mock.Anything).Return(nil)
		mockRedisOrder.On("SetJSON", mock.Anything, "COMPLETED", mock.Anything).Return(nil)

		err := orderUseCase.CreateOrder(orderDTO, "test-key")
		assert.NoError(t, err)
		mockRedisOrder.AssertExpectations(t)
		mockProductRepo.AssertExpectations(t)
		mockOrderRepo.AssertExpectations(t)
	})

	t.Run("order already completed", func(t *testing.T) {
		mockRedisOrder.On("KeyExists", mock.Anything).Return(true, nil)
		mockRedisOrder.On("GetOrderStatus", mock.Anything).Return("COMPLETED", nil)

		err := orderUseCase.CreateOrder(orderDTO, "test-key")
		assert.NoError(t, err)
		mockRedisOrder.AssertExpectations(t)
	})

}
