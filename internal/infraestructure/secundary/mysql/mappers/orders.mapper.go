package mappers

import (
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"

	"gorm.io/gorm"
)

func ToOrderDTO(order models.Order) dtos.OrderDTO {
	return dtos.OrderDTO{
		ID:           order.ID,
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
		CreatedAt:    order.CreatedAt,
		UpdatedAt:    order.UpdatedAt,
	}
}

func ToOrderModel(orderDTO dtos.OrderDTO) models.Order {
	return models.Order{
		Model: gorm.Model{
			ID:        orderDTO.ID,
			CreatedAt: orderDTO.CreatedAt,
			UpdatedAt: orderDTO.UpdatedAt,
		},
		CustomerName: orderDTO.CustomerName,
		TotalAmount:  orderDTO.TotalAmount,
	}
}

func ToOrderDTOList(orders []models.Order) []dtos.OrderDTO {
	orderDTOs := make([]dtos.OrderDTO, len(orders))
	for i, order := range orders {
		orderDTOs[i] = ToOrderDTO(order)
	}
	return orderDTOs
}

func ToOrderModelList(orderDTOs []dtos.OrderDTO) []models.Order {
	orders := make([]models.Order, len(orderDTOs))
	for i, orderDTO := range orderDTOs {
		orders[i] = ToOrderModel(orderDTO)
	}
	return orders
}
