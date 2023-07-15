package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ListProductsResponse struct {
	PaginatedResponse
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
func ListProductsController(c *gin.Context) {
	products := []model.Product{}

	response := NewPaginatedResponse(c)

	// Find products
	if err := db.Scopes(Paginate(products, &response)).Find(&products).Error; err != nil {
		log.Println(err)
		sendError(c, http.StatusInternalServerError, "failed to list products")
		return
	}

	response.Data = products

	c.JSON(http.StatusOK, response)
}
