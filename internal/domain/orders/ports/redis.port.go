package ports

import "velocity-technical-test/internal/domain/orders/dtos"

type IRedisOrder interface {
	SetJSON(key string, status string, value interface{}) error
	GetOrder(key string) (dtos.OrderDTORedis, error)
	GetOrderStatus(key string) (string, error)
	KeyExists(key string) (bool, error)
	UpdateOrderStatus(key string, newStatus string) error
}
