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

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the provided details
// @Tags orders
// @Accept json
// @Produce json
// @Param Idempotency-Key header string true "Idempotency Key"
// @Param order body request.OrderRequest true "Order Request"
// @Success 200 {object} response.BaseResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 500 {object} response.BaseResponse
// @Router /orders [post]
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
		if err.Error() == "409 Conflict: la clave de indempotencia ya existe con estado: IN_PROGRESS" {
			c.JSON(http.StatusConflict, response.BaseResponse{
				StatusCode: http.StatusConflict,
				Message:    err.Error(),
			})
			return
		}
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

// GetOrderWithItems godoc
// @Summary Get order with items
// @Description Get order details along with its items
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} response.BaseResponse
// @Failure 400 {object} response.BaseResponse
// @Failure 500 {object} response.BaseResponse
// @Router /orders/{id} [get]
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
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, response.BaseResponse{
				StatusCode: http.StatusNotFound,
				Message:    "Order not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.BaseResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Data:       mappers.ToOrderResponse(order),
	})
}
