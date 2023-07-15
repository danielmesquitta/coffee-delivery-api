package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/danielmesquitta/coffee-delivery-api/util"
	"github.com/gin-gonic/gin"
)

type CreateAddressRequest struct {
	ZipCode      string `json:"zipCode" validate:"required,len=9"`
	Street       string `json:"street" validate:"required"`
	Number       string `json:"number" validate:"required,numeric,gte=0"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required,len=2"`
}

type CreateUserRequest struct {
	PaymentMethod model.PaymentMethod  `json:"paymentMethod" validate:"required"`
	Address       CreateAddressRequest `json:"address" validate:"required"`
}

type CreateAddressResponse struct {
	ZipCode      string `json:"zipCode"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type CreateUserResponse struct {
	ID            uint                  `json:"id"`
	CreatedAt     time.Time             `json:"createdAt"`
	UpdatedAt     time.Time             `json:"updatedAt"`
	PaymentMethod model.PaymentMethod   `json:"paymentMethod"`
	AddressID     uint                  `json:"addressId"`
	Address       CreateAddressResponse `json:"address"`
}

// @BasePath /api/v1
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
func CreateUserController(ctx *gin.Context) {
	dto := CreateUserRequest{}

	ctx.ShouldBindJSON(&dto)

	log.Println(dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(ctx, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	user := model.User{
		PaymentMethod: dto.PaymentMethod,
		Address: model.Address{
			ZipCode:      dto.Address.ZipCode,
			Street:       dto.Address.Street,
			Number:       dto.Address.Number,
			Complement:   dto.Address.Complement,
			Neighborhood: dto.Address.Neighborhood,
			City:         dto.Address.City,
			State:        dto.Address.State,
		},
	}

	// Create user
	if err := db.Create(&user).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to create user")
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
