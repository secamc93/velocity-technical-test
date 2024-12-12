package repository

import (
	"sync"
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/domain/orders/ports"
	"velocity-technical-test/internal/infraestructure/secundary/mysql"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/mappers"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/models"
)

type Order struct {
	dbConnection mysql.DBConnection
}

var (
	instanceOrder *Order
	onceOrder     sync.Once
)

func NewOrder(db mysql.DBConnection) ports.IOrderRepository {
	onceOrder.Do(func() {
		instanceOrder = &Order{
			dbConnection: db,
		}
	})
	return instanceOrder
}

func (o *Order) CreateOrder(order dtos.OrderDTO) (uint, error) {
	db := o.dbConnection.GetDB()
	orderModel := models.Order{
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
	}
	err := db.Create(&orderModel).Error
	if err != nil {
		return 0, err
	}
	return orderModel.ID, nil
}

func (o *Order) CreateOrderItems(orderItems []dtos.OrderItemDTO) error {
	db := o.dbConnection.GetDB()
	for _, orderItem := range orderItems {
		orderItemModel := mappers.MapOrderItemDTOToModel(orderItem)
		err := db.Create(&orderItemModel).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *Order) GetOrderWithItems(orderID uint) (*dtos.OrderDTO, error) {
	db := o.dbConnection.GetDB()
	var orderModel models.Order
	err := db.Preload("OrderItems").
		Preload("OrderItems.Product").
		First(&orderModel, orderID).Error
	if err != nil {
		return nil, err
	}
	orderDTO := mappers.MapOrderModelToDTO(orderModel)
	return &orderDTO, nil
}
