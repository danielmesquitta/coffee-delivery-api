package controller

import (
	"net/http"
	"time"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ShowProductResponse struct {
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
// @Summary Show product
// @Description Show a job product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 200 {object} ShowProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product [get]
func ShowProductController(ctx *gin.Context) {
	// Get id from query and validate
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	product := model.Product{}

	// Find product
	if err := db.First(&product, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "product not found")
		return
	}

	ctx.JSON(http.StatusOK, product)
}
