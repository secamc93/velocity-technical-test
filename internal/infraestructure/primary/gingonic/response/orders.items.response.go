package response

type OrderItemResponse struct {
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Product   string  `json:"product"`
	Quantity  uint    `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderResponse struct {
	CustomerName string              `json:"customer_name"`
	TotalAmount  float64             `json:"total_amount"`
	Items        []OrderItemResponse `json:"items"`
}
