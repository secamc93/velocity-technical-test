package usecase

import (
	"fmt"
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/domain/orders/ports"
	producterrors "velocity-technical-test/internal/domain/products/errors"
	portsProduct "velocity-technical-test/internal/domain/products/ports"
	"velocity-technical-test/pkg/logger"
)

type IOrderUseCase interface {
	CreateOrder(orderDTO dtos.OrderDTO, indempotenciaKey string) error
	GetOrderWithItems(orderID uint) (*dtos.OrderDTO, error)
}

type Order struct {
	repoOrder   ports.IOrderRepository
	repoProduct portsProduct.IProductRepository
	redisOrder  ports.IRedisOrder
	logger      logger.ILogger
}

func NewOrder(repoOrder ports.IOrderRepository, repoProduct portsProduct.IProductRepository, redisOrder ports.IRedisOrder) IOrderUseCase {
	return &Order{
		repoOrder:   repoOrder,
		repoProduct: repoProduct,
		redisOrder:  redisOrder,
		logger:      logger.NewLogger(),
	}
}

func (o *Order) CreateOrder(orderDTO dtos.OrderDTO, indempotenciaKey string) error {
	// Validar existencia de la indempotenciaKey en Redis
	if err := o.checkIndempotenciaKey(indempotenciaKey); err != nil {
		if err.Error() == "Order already completed" {
			return nil
		}
		return err
	}

	// Crear un JSON inicial con estado IN_PROGRESS si no existe
	if err := o.createInitialOrderInRedis(indempotenciaKey, orderDTO); err != nil {
		return err
	}

	// Validar existencia y stock de productos, y calcular subtotal
	totalOrderPrice, err := o.validateAndCalculateOrder(orderDTO)
	if err != nil {
		return err
	}

	// Modificar el stock de productos después de la validación
	if err := o.updateProductStock(orderDTO); err != nil {
		return err
	}

	// Crear la orden y los items de la orden
	if err := o.createOrderAndItems(orderDTO, totalOrderPrice); err != nil {
		return err
	}

	// Actualizar el JSON en Redis con el status COMPLETED
	if err := o.redisOrder.SetJSON(indempotenciaKey, "COMPLETED", orderDTO); err != nil {
		o.logger.Error(err.Error())
		return err
	}

	return nil
}

func (o *Order) checkIndempotenciaKey(indempotenciaKey string) error {
	exists, err := o.redisOrder.KeyExists(indempotenciaKey)
	if err != nil {
		o.logger.Error(err.Error())
		return err
	}
	if exists {
		status, err := o.redisOrder.GetOrderStatus(indempotenciaKey)
		if err != nil {
			o.logger.Error(err.Error())
			return err
		}
		if status == "IN_PROGRESS" {
			err = fmt.Errorf("409 Conflict: la clave de indempotencia ya existe con estado: %s", status)
			o.logger.Error(err.Error())
			return err
		}
		if status == "COMPLETED" {
			o.logger.Info("Order already completed")
			return fmt.Errorf("Order already completed")
		}
	}
	return nil
}

func (o *Order) createInitialOrderInRedis(indempotenciaKey string, orderDTO dtos.OrderDTO) error {
	err := o.redisOrder.SetJSON(indempotenciaKey, "IN_PROGRESS", orderDTO)
	if err != nil {
		o.logger.Error(err.Error())
		return err
	}
	return nil
}

func (o *Order) validateAndCalculateOrder(orderDTO dtos.OrderDTO) (float64, error) {
	totalOrderPrice := 0.0
	for _, item := range orderDTO.Items {
		existProduct, err := o.repoProduct.ProductExists(item.ProductID)
		if err != nil {
			o.logger.Error(err.Error())
			return 0, err
		}
		if !existProduct {
			o.logger.Error(producterrors.ErrInvalidProduct.Error())
			return 0, producterrors.ErrInvalidProduct
		}

		stock, err := o.repoProduct.CountProductStock(item.ProductID)
		if err != nil {
			o.logger.Error(err.Error())
			return 0, err
		}

		if stock < item.Quantity {
			productName, err := o.repoProduct.GetProductName(item.ProductID)
			if err != nil {
				o.logger.Error(err.Error())
				return 0, err
			}
			err = fmt.Errorf("no hay stock suficiente para el producto: %s", productName)
			o.logger.Error(err.Error())
			return 0, err
		}

		price, err := o.repoProduct.GetProductPrice(item.ProductID)
		if err != nil {
			o.logger.Error(err.Error())
			return 0, err
		}

		item.Subtotal = price * float64(item.Quantity)
		totalOrderPrice += item.Subtotal
	}
	return totalOrderPrice, nil
}

func (o *Order) updateProductStock(orderDTO dtos.OrderDTO) error {
	for _, item := range orderDTO.Items {
		stock, err := o.repoProduct.CountProductStock(item.ProductID)
		if err != nil {
			o.logger.Error(err.Error())
			return err
		}

		newStock := stock - item.Quantity
		err = o.repoProduct.UpdateProductStock(item.ProductID, newStock)
		if err != nil {
			o.logger.Error(err.Error())
			return err
		}
	}
	return nil
}

func (o *Order) createOrderAndItems(orderDTO dtos.OrderDTO, totalOrderPrice float64) error {
	orderDTO.TotalAmount = totalOrderPrice
	orderID, err := o.repoOrder.CreateOrder(orderDTO)
	if err != nil {
		o.logger.Error(err.Error())
		return err
	}

	for i := range orderDTO.Items {
		orderDTO.Items[i].OrderID = orderID
	}

	err = o.repoOrder.CreateOrderItems(orderDTO.Items)
	if err != nil {
		o.logger.Error(err.Error())
		return err
	}
	return nil
}

func (o *Order) GetOrderWithItems(orderID uint) (*dtos.OrderDTO, error) {
	order, err := o.repoOrder.GetOrderWithItems(orderID)
	if err != nil {
		o.logger.Error(err.Error())
		return nil, err
	}

	return order, nil
}
