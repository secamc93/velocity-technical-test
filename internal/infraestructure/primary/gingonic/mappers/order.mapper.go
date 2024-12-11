package mappers

import (
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/request"
)

func ToOrderDTO(orderRequest request.OrderDTO) dtos.OrderDTO {
	return dtos.OrderDTO{
		ID:           orderRequest.ID,
		CustomerName: orderRequest.CustomerName,
		TotalAmount:  orderRequest.TotalAmount,
	}
}
