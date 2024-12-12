package mappers

import (
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"
)

func ToOrderDTO(order models.Order) dtos.OrderDTO {
	return dtos.OrderDTO{
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
	}
}

func ToOrderDTOList(orders []models.Order) []dtos.OrderDTO {
	orderDTOs := make([]dtos.OrderDTO, len(orders))
	for i, order := range orders {
		orderDTOs[i] = ToOrderDTO(order)
	}
	return orderDTOs
}

func MapOrderItemDTOToModel(orderItemDTO dtos.OrderItemDTO) models.OrderItem {
	return models.OrderItem{
		OrderID:   orderItemDTO.OrderID,
		ProductID: orderItemDTO.ProductID,
		Quantity:  orderItemDTO.Quantity,
		Subtotal:  orderItemDTO.Subtotal,
	}
}

func MapOrderModelToDTO(order models.Order) dtos.OrderDTO {
	return dtos.OrderDTO{
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
		Items:        MapOrderItemsModelToDTO(order.OrderItems),
	}
}

func MapOrderItemsModelToDTO(orderItems []models.OrderItem) []dtos.OrderItemDTO {
	orderItemsDTO := make([]dtos.OrderItemDTO, len(orderItems))
	for i, item := range orderItems {
		orderItemsDTO[i] = dtos.OrderItemDTO{
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Product:   item.Product.Name,
			Quantity:  item.Quantity,
			Subtotal:  item.Subtotal,
		}
	}
	return orderItemsDTO
}
