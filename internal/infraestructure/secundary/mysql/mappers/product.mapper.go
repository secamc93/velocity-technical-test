package mappers

import (
	"velocity-technical-test/internal/domain/products/dtos"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"

	"gorm.io/gorm"
)

func ToProductDTO(product models.Product) dtos.ProductDTO {
	return dtos.ProductDTO{
		ID:        product.ID,
		Name:      product.Name,
		Price:     product.Price,
		Stock:     uint(product.Stock),
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}

func ToProductModel(productDTO dtos.ProductDTO) models.Product {
	return models.Product{
		Model: gorm.Model{
			ID:        productDTO.ID,
			CreatedAt: productDTO.CreatedAt,
			UpdatedAt: productDTO.UpdatedAt,
		},
		Name:  productDTO.Name,
		Price: productDTO.Price,
		Stock: int(productDTO.Stock),
	}
}

func ToProductDTOList(products []models.Product) []dtos.ProductDTO {
	productDTOs := make([]dtos.ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = ToProductDTO(product)
	}
	return productDTOs
}

func ToProductModelList(productDTOs []dtos.ProductDTO) []models.Product {
	products := make([]models.Product, len(productDTOs))
	for i, productDTO := range productDTOs {
		products[i] = ToProductModel(productDTO)
	}
	return products
}
