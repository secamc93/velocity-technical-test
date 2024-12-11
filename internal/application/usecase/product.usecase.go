package usecase

import (
	"velocity-technical-test/internal/domain/products/dtos"
	"velocity-technical-test/internal/domain/products/ports"
)

type IProductUsecase interface {
	GetProducts() ([]dtos.ProductDTO, error)
}

type Product struct {
	repo ports.IProductRepository
}

func NewProduct(repo ports.IProductRepository) IProductUsecase {
	return &Product{
		repo: repo,
	}
}

func (p *Product) GetProducts() ([]dtos.ProductDTO, error) {
	return p.repo.GetProducts()
}
