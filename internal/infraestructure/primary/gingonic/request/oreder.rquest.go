package request

import "time"

type OrderItemRequest struct {
	ProductID uint    `json:"product_id"`
	Quantity  uint    `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}

type OrderRequest struct {
	CustomerName string             `json:"customer_name"`
	TotalAmount  float64            `json:"total_amount"`
	Items        []OrderItemRequest `json:"items"`
	CreatedAt    time.Time          `json:"created_at"`
	UdateAt      time.Time          `json:"update_at"`
}
