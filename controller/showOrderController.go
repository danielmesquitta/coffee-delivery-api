package controller

import (
	"net/http"
	"time"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ShowOrderResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Image       string    `json:"image"`
	Tags        []string  `json:"tags"`
}

// @BasePath /api/v1
// @Summary Show order
// @Description Show an order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id query string true "Order identification"
// @Success 200 {object} ShowOrderResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /order [get]
func ShowOrderController(c *gin.Context) {
	// Get id from query and validate
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "id is required")
		return
	}

	order := model.Order{}

	// Find order
	if err := db.Joins("User").First(&order, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "order not found")
		return
	}

	products := []model.Product{}

	// Find order products
	err := db.Model(&order).Association("Products").Find(&products)
	if err != nil {
		sendError(c, http.StatusNotFound, "failed to associate products to order")
		return
	}

	// Map products to orders
	for _, product := range products {
		order.Products = append(order.Products, &product)
	}

	c.JSON(http.StatusOK, order)
}
