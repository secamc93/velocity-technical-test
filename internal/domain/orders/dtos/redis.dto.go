package dtos

type Order struct {
	Status   string            `json:"status"`
	Response OrderItemDTORedis `json:"response"`
}

type OrderItemDTORedis struct {
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   string  `json:"product"`
	Quantity  uint    `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderDTORedis struct {
	CustomerName string              `json:"customer_name"`
	TotalAmount  float64             `json:"total_amount"`
	Items        []OrderItemDTORedis `json:"items"`
}
