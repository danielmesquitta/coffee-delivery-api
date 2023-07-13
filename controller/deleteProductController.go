package controller

import (
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

func DeleteProductController(ctx *gin.Context) {
	// Get id from query and validate
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	product := model.Product{}

	// Find product
	if err := db.First(&product, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	// Delete product
	if err := db.Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error updating product",
		})
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
