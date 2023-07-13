package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

func ListProductsController(ctx *gin.Context) {
	products := []model.Product{}

	// Find products
	if err := db.Find(&products).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to list products")
		return
	}

	ctx.JSON(http.StatusOK, products)
}
