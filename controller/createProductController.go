package controller

import (
	"log"

	"github.com/danielmesquitta/coffee-delivery-api/helper"
	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	dto := model.CreateProductDTO{}

	ctx.BindJSON(&dto)

	errs := helper.Validator.Validate(dto)

	if errs != nil {
		ctx.JSON(400, gin.H{
			"error": helper.Validator.FormatErrs(errs),
		})
		return
	}

	product := model.Product{
		Name:        dto.Name,
		Tags:        dto.Tags,
		Description: dto.Description,
		Price:       dto.Price,
	}

	if err := db.Create(&product); err != nil {
		log.Println("error: ", err)
		ctx.JSON(400, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	ctx.JSON(200, product)
}
