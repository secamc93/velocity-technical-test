package request

type OrderDTO struct {
	ID           uint    `json:"id"`
	CustomerName string  `json:"customer_name"`
	TotalAmount  float64 `json:"total_amount"`
}
