package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ListOrdersResponse struct {
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
func ListOrdersController(ctx *gin.Context) {
	orders := []model.Order{}

	// Find orders
	if err := db.Find(&orders).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to list orders")
		return
	}

	response := ListResponse{
		Data: orders,
	}

	ctx.JSON(http.StatusOK, response)
}
