package main

import (
	"github.com/gin-gonic/gin"
	handler "web-api/internal/core/adapters"
	"web-api/internal/core/adapters/repository"
	"web-api/internal/core/services"
)

var svc *services.ProductService

func main() {
	store := repository.NewProductPostgresRepository()
	svc = services.NewProductService(store)

	router := gin.Default()
	rhandler := handler.NewHTTPHandler(*svc)
	router.GET("/products", rhandler.GetAllProducts)
	err := router.Run(":3131")
	if err != nil {
		return
	}
}
