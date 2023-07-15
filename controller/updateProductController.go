package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/danielmesquitta/coffee-delivery-api/util"
	"github.com/gin-gonic/gin"
)

type UpdateProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price" validate:"numeric,gte=0"`
	Tags        []string `json:"tags"`
}

type UpdateProductResponse struct {
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
// @Summary Update product
// @Description Update a product
// @Tags Products
// @Accept json
// @Produce json
// @Param id query string true "Product Identification"
// @Param product body UpdateProductRequest true "Product data to Update"
// @Success 200 {object} UpdateProductResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [put]
func UpdateProductController(ctx *gin.Context) {
	// Get id from query and validate
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	dto := UpdateProductRequest{}

	ctx.BindJSON(&dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(ctx, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	product := model.Product{}

	// Find product
	if err := db.First(&product, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "product not found")
		return
	}

	product.Name = dto.Name
	product.Description = dto.Description
	product.Price = dto.Price
	product.Tags = dto.Tags

	// Save opening
	if err := db.Save(&product).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "error updating product")
		return
	}

	ctx.JSON(http.StatusOK, product)
}
