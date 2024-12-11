package usecase

import (
	"velocity-technical-test/internal/domain/orders/dtos"
	"velocity-technical-test/internal/domain/orders/ports"
	"velocity-technical-test/internal/domain/products/errors"
	portsProduct "velocity-technical-test/internal/domain/products/ports"
	"velocity-technical-test/pkg/logger"
)

type IOrderUseCase interface {
	CreateOrder(orderDTO dtos.OrderDTO) error
}

type Order struct {
	repoOrder   ports.IOrderRepository
	repoProduct portsProduct.IProductRepository
	logger      logger.ILogger
}

func NewOrder(repoOrder ports.IOrderRepository, repoProduct portsProduct.IProductRepository) IOrderUseCase {
	return &Order{
		repoOrder:   repoOrder,
		repoProduct: repoProduct,
		logger:      logger.NewLogger(),
	}
}

func (o *Order) CreateOrder(orderDTO dtos.OrderDTO) error {

	existProduct, err := o.repoProduct.ProductExists(orderDTO.ID)
	if err != nil || !existProduct {
		o.logger.Error(errors.ErrInvalidProduct.Error())
		return errors.ErrInvalidProduct
	}

	stock, err := o.repoProduct.CountProductStock(orderDTO.ID)
	if err != nil {
		o.logger.Error(errors.ErrInvalidProduct.Error())
		return errors.ErrNotStock
	}

	if stock < uint(orderDTO.TotalAmount) {
		o.logger.Error(errors.ErrNotStock.Error())
		return errors.ErrNotStock
	}

	return o.repoOrder.CreateOrder(orderDTO)
}
