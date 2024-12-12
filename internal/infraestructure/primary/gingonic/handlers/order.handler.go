package handlers

import (
	"net/http"
	"strconv"
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
	var orderRequest request.OrderRequest
	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	idempotencyKey := c.GetHeader("Idempotency-Key")
	if idempotencyKey == "" {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Idempotency-Key header is required",
		})
		return
	}

	orderDTO := mappers.ToOrderDTO(orderRequest)

	err := o.OrderHandler.CreateOrder(orderDTO, idempotencyKey)
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

func (o *OrderHandler) GetOrderWithItems(c *gin.Context) {
	orderID := c.Param("id")
	orderIDUint, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid order ID",
		})
		return
	}

	order, err := o.OrderHandler.GetOrderWithItems(uint(orderIDUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.BaseResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Data:       mappers.ToOrderResponse(order),
	})
}
