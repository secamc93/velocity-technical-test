package usecase

import (
	"velocity-technical-test/internal/domain/products/dtos"
	"velocity-technical-test/internal/domain/products/errors"
	"velocity-technical-test/internal/domain/products/ports"
	"velocity-technical-test/pkg/logger"
)

type IProductUsecase interface {
	GetProducts() ([]dtos.ProductDTO, error)
	UpdateProductStock(productID uint, newStock uint) error
}

type Product struct {
	repo   ports.IProductRepository
	logger logger.ILogger
}

func NewProduct(repo ports.IProductRepository) IProductUsecase {
	return &Product{
		repo:   repo,
		logger: logger.NewLogger(),
	}
}

func (p *Product) GetProducts() ([]dtos.ProductDTO, error) {
	return p.repo.GetProducts()
}

func (p *Product) UpdateProductStock(productID uint, newStock uint) error {
	existProduct, err := p.repo.ProductExists(productID)
	if err != nil || !existProduct {
		p.logger.Error(errors.ErrInvalidProduct.Error())
		return errors.ErrInvalidProduct
	}

	return p.repo.UpdateProductStock(productID, newStock)
}
