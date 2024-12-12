package mappers

import (
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/request"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/response"
)

func ToOrderDTO(orderRequest request.OrderRequest) dtos.OrderDTO {
	items := make([]dtos.OrderItemDTO, len(orderRequest.Items))
	for i, item := range orderRequest.Items {
		items[i] = dtos.OrderItemDTO{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Subtotal:  item.Subtotal,
		}
	}
	return dtos.OrderDTO{
		CustomerName: orderRequest.CustomerName,
		TotalAmount:  orderRequest.TotalAmount,
		Items:        items,
	}
}

func ToOrderResponse(orderDTO *dtos.OrderDTO) response.OrderResponse {
	items := make([]response.OrderItemResponse, len(orderDTO.Items))
	for i, item := range orderDTO.Items {
		items[i] = response.OrderItemResponse{
			OrderID:   item.OrderID,
			ProductID: item.ProductID,
			Product:   item.Product,
			Quantity:  item.Quantity,
			Subtotal:  item.Subtotal,
		}
	}
	return response.OrderResponse{
		CustomerName: orderDTO.CustomerName,
		TotalAmount:  orderDTO.TotalAmount,
		Items:        items,
	}
}
