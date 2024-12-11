package handlers

import (
	"net/http"
	"velocity-technical-test/internal/application/usecase"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/mappers"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/request"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/response"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	OrderHandler usecase.IOrderUseCase
}

func NewOrderHandler(orderHandler usecase.IOrderUseCase) *OrderHandler {
	return &OrderHandler{
		OrderHandler: orderHandler,
	}
}

func (o *OrderHandler) CreateOrder(c *gin.Context) {
	var orderRequest request.OrderDTO
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	orderDTO := mappers.ToOrderDTO(orderRequest)

	err := o.OrderHandler.CreateOrder(orderDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "Order created successfully",
	})
}
