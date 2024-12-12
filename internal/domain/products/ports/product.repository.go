package ports

import "velocity-technical-test/internal/domain/products/dtos"

type IProductRepository interface {
	GetProducts() ([]dtos.ProductDTO, error)
	UpdateProductStock(productID uint, newStock uint) error
	ProductExists(productID uint) (bool, error)
	CountProductStock(productID uint) (uint, error)
	GetProductName(productID uint) (string, error)
	GetProductPrice(productID uint) (float64, error)
}
