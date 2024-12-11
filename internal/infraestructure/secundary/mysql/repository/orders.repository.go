package repository

import (
	"sync"
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/domain/orders/ports"
	"velocity-technical-test/internal/infraestructure/secundary/mysql"
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

func (o *Order) CreateOrder(order dtos.OrderDTO) error {
	db := o.dbConnection.GetDB()
	orderModel := models.Order{
		CustomerName: order.CustomerName,
		TotalAmount:  order.TotalAmount,
	}
	err := db.Create(&orderModel).Error
	if err != nil {
		return err
	}
	return nil
}
