package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

func DeleteProductController(ctx *gin.Context) {
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

	// Delete product
	if err := db.Delete(&product).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to delete product")
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
