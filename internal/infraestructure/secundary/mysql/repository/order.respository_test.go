package repository

import (
	"testing"
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"
	"velocity-technical-test/mocks"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateOrder(t *testing.T) {
	// Crear DB en memoria
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)

	// Migrar los modelos necesarios
	err = db.AutoMigrate(&models.Order{})
	require.NoError(t, err)

	// Crear el mock de DBConnection
	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	orderRepo := NewOrder(mockDB)

	// Crear un OrderDTO de prueba
	orderDTO := dtos.OrderDTO{
		CustomerName: "John Doe",
		TotalAmount:  500.0,
		// OrderItems se crean después en otros tests
	}

	// Ejecutar el método
	orderID, err := orderRepo.CreateOrder(orderDTO)
	require.NoError(t, err)
	require.NotZero(t, orderID)

	// Verificar que el pedido se creó en la BD
	var orderModel models.Order
	err = db.First(&orderModel, orderID).Error
	require.NoError(t, err)
	require.Equal(t, "John Doe", orderModel.CustomerName)
	require.Equal(t, 500.0, orderModel.TotalAmount)

	mockDB.AssertExpectations(t)
}

func TestCreateOrderItems(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)

	// Migrar modelos
	// Asumimos que Order tiene One-to-Many con OrderItems, y OrderItem tiene una FK a Product
	err = db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Product{})
	require.NoError(t, err)

	// Insertar un producto de prueba
	product := models.Product{Name: "Product A", Price: 100.0, Stock: 10}
	err = db.Create(&product).Error
	require.NoError(t, err)

	// Insertar un order para asociarle items
	order := models.Order{CustomerName: "Jane Smith", TotalAmount: 200.0}
	err = db.Create(&order).Error
	require.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	orderRepo := NewOrder(mockDB)

	// Crear OrderItemDTOs
	orderItemsDTO := []dtos.OrderItemDTO{
		{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  2,
			Subtotal:  product.Price * 2,
		},
		{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  1,
			Subtotal:  product.Price,
		},
	}

	// Ejecutar método
	err = orderRepo.CreateOrderItems(orderItemsDTO)
	require.NoError(t, err)

	// Verificar en la DB que se crearon los items
	var orderItems []models.OrderItem
	err = db.Where("order_id = ?", order.ID).Find(&orderItems).Error
	require.NoError(t, err)
	require.Len(t, orderItems, 2)

	// Validar que los datos coincidan
	require.Equal(t, product.ID, orderItems[0].ProductID)
	require.Equal(t, uint(2), orderItems[0].Quantity)
	require.Equal(t, product.Price*2, orderItems[0].Subtotal)

	require.Equal(t, product.ID, orderItems[1].ProductID)
	require.Equal(t, uint(1), orderItems[1].Quantity)
	require.Equal(t, product.Price, orderItems[1].Subtotal)

	mockDB.AssertExpectations(t)
}

func TestGetOrderWithItems(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)

	// Migrar modelos
	err = db.AutoMigrate(&models.Order{}, &models.OrderItem{}, &models.Product{})
	require.NoError(t, err)

	// Crear datos de prueba
	product := models.Product{Name: "Product B", Price: 150.0, Stock: 5}
	err = db.Create(&product).Error
	require.NoError(t, err)

	order := models.Order{CustomerName: "Alice", TotalAmount: 450.0}
	err = db.Create(&order).Error
	require.NoError(t, err)

	orderItems := []models.OrderItem{
		{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  2,
			Subtotal:  150.0 * 2,
		},
		{
			OrderID:   order.ID,
			ProductID: product.ID,
			Quantity:  1,
			Subtotal:  150.0,
		},
	}
	err = db.Create(&orderItems).Error
	require.NoError(t, err)

	mockDB := new(mocks.DBConnection)
	mockDB.On("GetDB").Return(db)

	orderRepo := NewOrder(mockDB)

	// Ejecutar el método
	orderDTO, err := orderRepo.GetOrderWithItems(order.ID)
	require.NoError(t, err)
	require.NotNil(t, orderDTO)

	// Validar datos básicos del pedido
	require.Equal(t, "Alice", orderDTO.CustomerName)
	require.Equal(t, 450.0, orderDTO.TotalAmount)
	require.Len(t, orderDTO.Items, 2)

	// Validar datos del primer item
	require.Equal(t, product.ID, orderDTO.Items[0].ProductID)
	require.Equal(t, uint(2), orderDTO.Items[0].Quantity)
	require.Equal(t, 150.0*2, orderDTO.Items[0].Subtotal)

	// Validar datos del segundo item
	require.Equal(t, product.ID, orderDTO.Items[1].ProductID)
	require.Equal(t, uint(1), orderDTO.Items[1].Quantity)
	require.Equal(t, 150.0, orderDTO.Items[1].Subtotal)

	mockDB.AssertExpectations(t)
}
