package dtos

import "time"

type OrderDTO struct {
	ID           uint
	CustomerName string
	TotalAmount  float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
