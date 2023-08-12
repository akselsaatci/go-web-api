package handler

import (
	"web-api/internal/core/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	svc services.ProductService
}

func NewHTTPHandler(ProductService services.ProductService) *HTTPHandler {
	return &HTTPHandler{
		svc: ProductService,
	}
}

func (h *HTTPHandler) GetAllProducts(ctx *gin.Context) {

	products, err := h.svc.GetAllProducts()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, products)
}
