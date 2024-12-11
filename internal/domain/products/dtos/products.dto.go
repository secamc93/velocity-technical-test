package dtos

import "time"

type ProductDTO struct {
	ID        uint
	Name      string
	Price     float64
	Stock     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
