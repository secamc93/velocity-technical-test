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

	urlBase := os.Getenv("URL_BASE")
	if urlBase == "" {
		urlBase = "/api"
	}

	api := router.Group(urlBase)
	{
		handler := handlers.NewProductHandler(productUseCase)
		api.GET("/products", handler.GetProducts)
	}

	return router
}
