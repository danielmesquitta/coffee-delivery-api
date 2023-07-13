package controller

import (
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/helper"
	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

func UpdateProductController(ctx *gin.Context) {
	// Get id from query and validate
	id := ctx.Query("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	dto := model.UpdateProductDTO{}

	ctx.BindJSON(&dto)

	// Validate DTO
	errs := helper.Validator.Validate(dto)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": helper.Validator.FormatErrs(errs),
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

	product.Name = dto.Name
	product.Description = dto.Description
	product.Price = dto.Price
	product.Tags = dto.Tags

	// Save opening
	if err := db.Save(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error updating product",
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
