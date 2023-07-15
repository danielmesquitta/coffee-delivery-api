package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

// @Summary Delete product
// @Description Delete a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product identification"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /product [delete]
func DeleteProductController(c *gin.Context) {
	// Get id from query and validate
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "id is required")
		return
	}

	product := model.Product{}

	// Find product
	if err := db.First(&product, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "product not found")
		return
	}

	// Delete product
	if err := db.Delete(&product).Error; err != nil {
		log.Println(err)
		sendError(c, http.StatusInternalServerError, "failed to delete product")
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
