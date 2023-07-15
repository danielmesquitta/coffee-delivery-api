package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/danielmesquitta/coffee-delivery-api/util"
	"github.com/gin-gonic/gin"
)

type UpdateAddressRequest struct {
	ZipCode      string `json:"zipCode" validate:"required,len=9"`
	Street       string `json:"street" validate:"required"`
	Number       string `json:"number" validate:"required,numeric,gte=0"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	City         string `json:"city" validate:"required"`
	State        string `json:"state" validate:"required,len=2"`
}

type UpdateUserRequest struct {
	PaymentMethod model.PaymentMethod  `json:"paymentMethod" validate:"required"`
	Address       UpdateAddressRequest `json:"address" validate:"required"`
}

type UpdateAddressResponse struct {
	ID           uint      `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	ZipCode      string    `json:"zipCode"`
	Street       string    `json:"street"`
	Number       string    `json:"number"`
	Complement   string    `json:"complement"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	State        string    `json:"state"`
}

// @BasePath /api/v1
// @Summary Update user
// @Description Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User Identification"
// @Param user body UpdateUserRequest true "User data to Update"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [put]
func UpdateUserController(c *gin.Context) {
	// Get id from query and validate
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, "id is required")
		return
	}

	dto := UpdateUserRequest{}

	c.BindJSON(&dto)

	// Validate DTO
	errs := util.Validator.Validate(dto)
	if errs != nil {
		sendError(c, http.StatusBadRequest, util.Validator.FormatErrs(errs))
		return
	}

	user := model.User{}

	// Find user
	if err := db.First(&user, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "user not found")
		return
	}

	user.PaymentMethod = dto.PaymentMethod
	user.Address = &model.Address{
		ZipCode:      dto.Address.ZipCode,
		Street:       dto.Address.Street,
		Number:       dto.Address.Number,
		Complement:   dto.Address.Complement,
		Neighborhood: dto.Address.Neighborhood,
		City:         dto.Address.City,
		State:        dto.Address.State,
	}

	// Save opening
	if err := db.Save(&user).Error; err != nil {
		log.Println(err)
		sendError(c, http.StatusInternalServerError, "error updating user")
		return
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
