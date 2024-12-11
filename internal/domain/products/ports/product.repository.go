package ports

import "velocity-technical-test/internal/domain/products/dtos"

type IProductRepository interface {
	GetProducts() ([]dtos.ProductDTO, error)
}
