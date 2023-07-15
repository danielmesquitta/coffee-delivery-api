package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ListProductsResponse struct {
	Data []ShowProductResponse `json:"data"`
}

// @BasePath /api/v1
// @Summary List products
// @Description List all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} ListProductsResponse
// @Failure 500 {object} ErrorResponse
// @Router /products [get]
func ListProductsController(ctx *gin.Context) {
	products := []model.Product{}

	// Find products
	if err := db.Find(&products).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to list products")
		return
	}

	response := struct {
		Data []model.Product
	}{
		Data: products,
	}

	ctx.JSON(http.StatusOK, response)
}
