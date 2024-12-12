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
	instanceProduct *Product
	onceProduct     sync.Once
)

func NewProduct(db mysql.DBConnection) ports.IProductRepository {
	onceProduct.Do(func() {
		instanceProduct = &Product{
			dbConnection: db,
		}
	})
	return instanceProduct
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

func (r *Product) UpdateProductStock(productID uint, newStock uint) error {
	db := r.dbConnection.GetDB()
	err := db.Debug().Model(&models.Product{}).Where("id = ?", productID).Update("stock", newStock).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Product) ProductExists(productID uint) (bool, error) {
	var count int64
	db := r.dbConnection.GetDB()
	err := db.Model(&models.Product{}).Where("id = ?", productID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Product) CountProductStock(productID uint) (uint, error) {
	var stock uint
	db := r.dbConnection.GetDB()
	err := db.Model(&models.Product{}).Where("id = ?", productID).Pluck("stock", &stock).Error
	if err != nil {
		return 0, err
	}
	return stock, nil
}

func (r *Product) GetProductName(productID uint) (string, error) {
	var productName string
	db := r.dbConnection.GetDB()
	err := db.Model(&models.Product{}).Where("id = ?", productID).Pluck("name", &productName).Error
	if err != nil {
		return "", err
	}
	return productName, nil
}

func (r *Product) GetProductPrice(productID uint) (float64, error) {
	var price float64
	db := r.dbConnection.GetDB()
	err := db.Model(&models.Product{}).Where("id = ?", productID).Pluck("price", &price).Error
	if err != nil {
		return 0, err
	}
	return price, nil
}
