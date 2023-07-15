package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

// @Summary Delete user
// @Description Delete an user
// @Tags Users
// @Accept json
// @Produce json
// @Param id query string true "User identification"
// @Success 204
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /user [delete]
func DeleteUserController(ctx *gin.Context) {
	// Get id from query and validate
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, "id is required")
		return
	}

	user := model.User{}

	// Find user
	if err := db.First(&user, id).Error; err != nil {
		log.Println(err.Error())
		sendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	address := model.Address{
		ID: user.AddressID,
	}

	// Delete user address
	if err := db.Delete(&address).Error; err != nil {
		log.Println(err.Error())
		sendError(ctx, http.StatusInternalServerError, "failed to delete user address")
		return
	}

	// Delete user
	if err := db.Delete(&user).Error; err != nil {
		log.Println(err.Error())
		sendError(ctx, http.StatusInternalServerError, "failed to delete user")
		return
	}

	ctx.Writer.WriteHeader(http.StatusNoContent)
}
