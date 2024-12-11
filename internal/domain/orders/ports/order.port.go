package ports

import "velocity-technical-test/internal/domain/orders/dtos"

type IOrderRepository interface {
	CreateOrder(order dtos.OrderDTO) error
}
