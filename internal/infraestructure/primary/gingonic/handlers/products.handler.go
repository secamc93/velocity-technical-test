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

type ProductHandler struct {
	ProductUsecase usecase.IProductUsecase
}

func NewProductHandler(productUsecase usecase.IProductUsecase) *ProductHandler {
	return &ProductHandler{
		ProductUsecase: productUsecase,
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.ProductUsecase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	productResponse := mappers.ToProductResponseList(products)
	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       productResponse,
	})
}

func (h *ProductHandler) UpdateProductStock(c *gin.Context) {
	request := request.StockRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productIDStr := c.Param("id")
	productID, err := strconv.ParseUint(productIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.ProductUsecase.UpdateProductStock(uint(productID), request.NewStock)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		StatusCode: http.StatusOK,
		Message:    "Stock updated successfully",
	})
}
