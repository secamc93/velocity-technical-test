package repository

import (
	"sync"
	"velocity-technical-test/internal/domain/products/dtos"
	"velocity-technical-test/internal/domain/products/ports"
	"velocity-technical-test/internal/infraestructure/secundary/mysql"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/mappers"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"
)

type Product struct {
	dbConnection mysql.DBConnection
}

var (
	instance *Product
	once     sync.Once
)

func NewProduct(db mysql.DBConnection) ports.IProductRepository {
	once.Do(func() {
		instance = &Product{
			dbConnection: db,
		}
	})
	return instance
}

func (r *Product) GetProducts() ([]dtos.ProductDTO, error) {
	var productModels []models.Product
	db := r.dbConnection.GetDB()
	err := db.Model(&productModels).Find(&productModels).Error
	if err != nil {
		return nil, err
	}

	productDTOS := mappers.ToProductDTOList(productModels)

	return productDTOS, nil
}
