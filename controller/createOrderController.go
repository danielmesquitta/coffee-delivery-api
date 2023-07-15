package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/danielmesquitta/coffee-delivery-api/util"
	"github.com/gin-gonic/gin"
)

type AssociateProductRequest struct {
	ID uint `json:"id" validate:"required"`
}

type CreateOrderRequest struct {
	Products []AssociateProductRequest `json:"products"`
	UserID   uint                      `json:"userId"`
}

// @BasePath /api/v1
// @Summary Create order
// @Description Create a new order
// @Tags Orders
// @Accept json
// @Produce json
// @Param request body CreateOrderRequest true "Request body"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /order [post]
func CreateOrderController(ctx *gin.Context) {
	dto := CreateOrderRequest{}

	ctx.ShouldBindJSON(&dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(ctx, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	order := model.Order{
		UserID: dto.UserID,
	}

	// Associate products
	for _, product := range dto.Products {
		order.Products = append(order.Products, &model.Product{
			ID: product.ID,
		})
	}

	// Create order
	if err := db.Create(&order).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to create order")
		return
	}

	ctx.Writer.WriteHeader(http.StatusCreated)
}
