package usecase

import (
	"errors"
	"testing"
	"velocity-technical-test/internal/domain/products/dtos"
	producterrors "velocity-technical-test/internal/domain/products/errors"
	"velocity-technical-test/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	mockRepo.On("GetProducts").Return([]dtos.ProductDTO{{ID: 1, Name: "Product1"}}, nil)

	productUsecase := NewProduct(mockRepo)

	products, err := productUsecase.GetProducts()

	assert.NoError(t, err)
	assert.Len(t, products, 1)
	assert.Equal(t, uint(1), products[0].ID)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProductStock_Success(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	mockRepo.On("ProductExists", uint(1)).Return(true, nil)
	mockRepo.On("UpdateProductStock", uint(1), uint(10)).Return(nil)

	productUsecase := NewProduct(mockRepo)

	err := productUsecase.UpdateProductStock(1, 10)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProductStock_ProductNotExist(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	mockRepo.On("ProductExists", uint(1)).Return(false, nil)

	productUsecase := NewProduct(mockRepo)

	err := productUsecase.UpdateProductStock(1, 10)

	assert.Error(t, err)
	assert.Equal(t, producterrors.ErrInvalidProduct, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateProductStock_ErrorCheckingExistence(t *testing.T) {
	mockRepo := new(mocks.IProductRepository)
	mockRepo.On("ProductExists", uint(1)).Return(false, errors.New("some error"))

	productUsecase := NewProduct(mockRepo)

	err := productUsecase.UpdateProductStock(1, 10)

	assert.Error(t, err)
	assert.Equal(t, producterrors.ErrInvalidProduct, err)
	mockRepo.AssertExpectations(t)
}
