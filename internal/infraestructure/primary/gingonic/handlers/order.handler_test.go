package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"velocity-technical-test/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOrder_Success(t *testing.T) {
	mockOrderUsecase := new(mocks.IOrderUseCase)
	mockOrderUsecase.On("CreateOrder", mock.Anything, "test-key").Return(nil)

	orderHandler := NewOrderHandler(mockOrderUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)

	orderJSON := `{"customer_id": 1, "items": [{"product_id": 1, "quantity": 2}]}`
	req, _ := http.NewRequest(http.MethodPost, "/orders", strings.NewReader(orderJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", "test-key")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockOrderUsecase.AssertExpectations(t)
}

func TestCreateOrder_MissingIdempotencyKey(t *testing.T) {
	mockOrderUsecase := new(mocks.IOrderUseCase)

	orderHandler := NewOrderHandler(mockOrderUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)

	orderJSON := `{"customer_id": 1, "items": [{"product_id": 1, "quantity": 2}]}`
	req, _ := http.NewRequest(http.MethodPost, "/orders", strings.NewReader(orderJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockOrderUsecase.AssertExpectations(t)
}

func TestCreateOrder_Conflict(t *testing.T) {
	mockOrderUsecase := new(mocks.IOrderUseCase)
	mockOrderUsecase.On("CreateOrder", mock.Anything, "test-key").Return(errors.New("409 Conflict: la clave de indempotencia ya existe con estado: IN_PROGRESS"))

	orderHandler := NewOrderHandler(mockOrderUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/orders", orderHandler.CreateOrder)

	orderJSON := `{"customer_id": 1, "items": [{"product_id": 1, "quantity": 2}]}`
	req, _ := http.NewRequest(http.MethodPost, "/orders", strings.NewReader(orderJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Idempotency-Key", "test-key")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	mockOrderUsecase.AssertExpectations(t)
}

func TestGetOrderWithItems_NotFound(t *testing.T) {
	mockOrderUsecase := new(mocks.IOrderUseCase)
	mockOrderUsecase.On("GetOrderWithItems", uint(1)).Return(nil, errors.New("record not found"))

	orderHandler := NewOrderHandler(mockOrderUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/orders/:id", orderHandler.GetOrderWithItems)

	req, _ := http.NewRequest(http.MethodGet, "/orders/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	mockOrderUsecase.AssertExpectations(t)
}

func TestGetOrderWithItems_InvalidID(t *testing.T) {
	mockOrderUsecase := new(mocks.IOrderUseCase)

	orderHandler := NewOrderHandler(mockOrderUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/orders/:id", orderHandler.GetOrderWithItems)

	req, _ := http.NewRequest(http.MethodGet, "/orders/invalid", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockOrderUsecase.AssertExpectations(t)
}
