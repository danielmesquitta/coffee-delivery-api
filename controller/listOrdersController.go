package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ListOrdersResponse struct {
	PaginatedResponse
	Data []ShowOrderResponse `json:"data"`
}

// @BasePath /api/v1
// @Summary List orders
// @Description List all orders
// @Tags Orders
// @Accept json
// @Produce json
// @Success 200 {object} ListOrdersResponse
// @Failure 500 {object} ErrorResponse
// @Router /orders [get]
func ListOrdersController(c *gin.Context) {
	orders := []model.Order{}

	response := NewPaginatedResponse(c)

	// Find orders
	if err := db.Scopes(Paginate(orders, &response)).Find(&orders).Error; err != nil {
		log.Println(err)
		sendError(c, http.StatusInternalServerError, "failed to list orders")
		return
	}

	response.Data = orders

	c.JSON(http.StatusOK, response)
}
