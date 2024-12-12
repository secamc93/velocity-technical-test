package ports

import "velocity-technical-test/internal/domain/orders/dtos"

type IOrderRepository interface {
	CreateOrder(order dtos.OrderDTO) (uint, error)
	CreateOrderItems(orderItems []dtos.OrderItemDTO) error
	GetOrderWithItems(orderID uint) (*dtos.OrderDTO, error)
}
