package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"velocity-technical-test/internal/domain/products/dtos"
	"velocity-technical-test/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts_Success(t *testing.T) {
	mockProductUsecase := new(mocks.IProductUsecase)
	mockProductUsecase.On("GetProducts").Return([]dtos.ProductDTO{{ID: 1, Name: "Product1"}}, nil)

	productHandler := NewProductHandler(mockProductUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/products", productHandler.GetProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockProductUsecase.AssertExpectations(t)
}

func TestGetProducts_Error(t *testing.T) {
	mockProductUsecase := new(mocks.IProductUsecase)
	mockProductUsecase.On("GetProducts").Return(nil, errors.New("some error"))

	productHandler := NewProductHandler(mockProductUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/products", productHandler.GetProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockProductUsecase.AssertExpectations(t)
}

func TestUpdateProductStock_Success(t *testing.T) {
	mockProductUsecase := new(mocks.IProductUsecase)
	mockProductUsecase.On("UpdateProductStock", uint(1), uint(10)).Return(nil)

	productHandler := NewProductHandler(mockProductUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/products/:id/stock", productHandler.UpdateProductStock)

	stockJSON := `{"new_stock": 10}`
	req, _ := http.NewRequest(http.MethodPut, "/products/1/stock", strings.NewReader(stockJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockProductUsecase.AssertExpectations(t)
}

func TestUpdateProductStock_InvalidProductID(t *testing.T) {
	mockProductUsecase := new(mocks.IProductUsecase)

	productHandler := NewProductHandler(mockProductUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/products/:id/stock", productHandler.UpdateProductStock)

	stockJSON := `{"new_stock": 10}`
	req, _ := http.NewRequest(http.MethodPut, "/products/invalid/stock", strings.NewReader(stockJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockProductUsecase.AssertExpectations(t)
}

func TestUpdateProductStock_BindJSONError(t *testing.T) {
	mockProductUsecase := new(mocks.IProductUsecase)

	productHandler := NewProductHandler(mockProductUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/products/:id/stock", productHandler.UpdateProductStock)

	invalidJSON := `{"new_stock": "invalid"}`
	req, _ := http.NewRequest(http.MethodPut, "/products/1/stock", strings.NewReader(invalidJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockProductUsecase.AssertExpectations(t)
}

func TestUpdateProductStock_Error(t *testing.T) {
	mockProductUsecase := new(mocks.IProductUsecase)
	mockProductUsecase.On("UpdateProductStock", uint(1), uint(10)).Return(errors.New("some error"))

	productHandler := NewProductHandler(mockProductUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/products/:id/stock", productHandler.UpdateProductStock)

	stockJSON := `{"new_stock": 10}`
	req, _ := http.NewRequest(http.MethodPut, "/products/1/stock", strings.NewReader(stockJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockProductUsecase.AssertExpectations(t)
}
