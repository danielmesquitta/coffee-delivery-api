package controller

import (
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/helper"
	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

func CreateProductController(ctx *gin.Context) {
	dto := model.CreateProductDTO{}

	ctx.ShouldBindJSON(&dto)

	// Validate DTO
	errs := helper.Validator.Validate(dto)
	if errs != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": helper.Validator.FormatErrs(errs),
		})
		return
	}

	product := model.Product{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Tags:        dto.Tags,
	}

	// Create product
	if err := db.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
