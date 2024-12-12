package dtos

type OrderItemDTO struct {
	OrderID   uint
	ProductID uint
	Product   string
	Quantity  uint
	Subtotal  float64
}

type OrderDTO struct {
	CustomerName string
	TotalAmount  float64
	Items        []OrderItemDTO
}
