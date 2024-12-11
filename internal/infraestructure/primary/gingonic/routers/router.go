package routers

import (
	"os"
	"velocity-technical-test/internal/application/usecase"
	"velocity-technical-test/internal/infraestructure/primary/gingonic/handlers"
	"velocity-technical-test/internal/infraestructure/secundary/mysql"
	"velocity-technical-test/internal/infraestructure/secundary/mysql/repository"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	dbConnection := mysql.NewDBConnection()
	productRepo := repository.NewProduct(dbConnection)
	productUseCase := usecase.NewProduct(productRepo)
	orderRepo := repository.NewOrder(dbConnection)
	ordersUseCase := usecase.NewOrder(orderRepo, productRepo)

	urlBase := os.Getenv("URL_BASE")
	if urlBase == "" {
		urlBase = "/api"
	}

	api := router.Group(urlBase)
	{
		handlerProduct := handlers.NewProductHandler(productUseCase)
		handlerOrder := handlers.NewOrderHandler(ordersUseCase)
		api.GET("/products", handlerProduct.GetProducts)
		api.PUT("/products/:id/stock", handlerProduct.UpdateProductStock)
		api.POST("/orders", handlerOrder.CreateOrder)
	}

	return router
}
