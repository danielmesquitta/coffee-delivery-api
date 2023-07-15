package controller

import (
	"log"
	"net/http"

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

// @BasePath /api/v1
// @Summary Create user
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 201
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
func CreateUserController(c *gin.Context) {
	dto := CreateUserRequest{}

	c.ShouldBindJSON(&dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(c, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	user := model.User{
		PaymentMethod: dto.PaymentMethod,
		Address: &model.Address{
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
		sendError(c, http.StatusInternalServerError, "failed to create user")
		return
	}

	c.Writer.WriteHeader(http.StatusCreated)
}
