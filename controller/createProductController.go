package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/danielmesquitta/coffee-delivery-api/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type CreateProductRequest struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Price       float64        `json:"price" validate:"required,numeric,gte=0"`
	Tags        pq.StringArray `json:"tags"`
}

type CreateProductResponse struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Image       string         `json:"image"`
	Tags        pq.StringArray `json:"tags"`
}

func CreateProductController(ctx *gin.Context) {
	dto := CreateProductRequest{}

	ctx.ShouldBindJSON(&dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(ctx, http.StatusBadRequest, util.Validator.FormatErrs(errs))
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
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to create product")
		return
	}

	ctx.JSON(http.StatusCreated, product)
}
