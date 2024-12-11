package handlers

import (
	"net/http"
	"velocity-technical-test/internal/application/usecase"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/mappers"
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
