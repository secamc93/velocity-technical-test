package repository

import (
	"testing"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"
	"velocity-technical-test/mocks"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetProducts(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Product{})
	assert.NoError(t, err)

	err = db.Create(&models.Product{Name: "Test Product"}).Error
	assert.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	productRepo := NewProduct(mockDB)

	products, err := productRepo.GetProducts()

	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Len(t, products, 1)
	assert.Equal(t, "Test Product", products[0].Name)

	mockDB.AssertExpectations(t)
}

func TestUpdateProductStock(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Product{})
	assert.NoError(t, err)

	product := models.Product{Name: "Producto A", Stock: 10}
	err = db.Create(&product).Error
	assert.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	productRepo := NewProduct(mockDB)

	newStock := uint(50)
	err = productRepo.UpdateProductStock(product.ID, newStock)
	assert.NoError(t, err)

	var updatedProduct models.Product
	err = db.First(&updatedProduct, product.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, int(newStock), updatedProduct.Stock)

	mockDB.AssertExpectations(t)
}

func TestProductExists(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Product{})
	assert.NoError(t, err)

	product := models.Product{Name: "Producto A", Stock: 10}
	err = db.Create(&product).Error
	assert.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	productRepo := NewProduct(mockDB)

	exists, err := productRepo.ProductExists(product.ID)
	assert.NoError(t, err)
	assert.True(t, exists)

	nonExistentID := uint(9999)
	exists, err = productRepo.ProductExists(nonExistentID)
	assert.NoError(t, err)
	assert.False(t, exists)

	mockDB.AssertExpectations(t)
}

func TestCountProductStock(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Product{})
	assert.NoError(t, err)

	product := models.Product{Name: "Producto A", Stock: 20, Price: 100.0}
	err = db.Create(&product).Error
	assert.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	productRepo := NewProduct(mockDB)

	stock, err := productRepo.CountProductStock(product.ID)
	assert.NoError(t, err)
	assert.Equal(t, uint(20), stock)

	stock, err = productRepo.CountProductStock(9999)
	assert.NoError(t, err)
	assert.Equal(t, uint(0), stock)

	mockDB.AssertExpectations(t)
}

func TestGetProductName(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Product{})
	assert.NoError(t, err)

	product := models.Product{Name: "Producto B", Stock: 10, Price: 200.0}
	err = db.Create(&product).Error
	assert.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	productRepo := NewProduct(mockDB)

	name, err := productRepo.GetProductName(product.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Producto B", name)

	name, err = productRepo.GetProductName(9999)
	assert.NoError(t, err)
	assert.Equal(t, "", name)

	mockDB.AssertExpectations(t)
}

func TestGetProductPrice(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&models.Product{})
	assert.NoError(t, err)

	product := models.Product{Name: "Producto C", Stock: 5, Price: 300.0}
	err = db.Create(&product).Error
	assert.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	productRepo := NewProduct(mockDB)

	price, err := productRepo.GetProductPrice(product.ID)
	assert.NoError(t, err)
	assert.Equal(t, 300.0, price)

	price, err = productRepo.GetProductPrice(9999)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, price)

	mockDB.AssertExpectations(t)
}
