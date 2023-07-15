package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ListUsersResponse struct {
	Data []ShowUserResponse `json:"data"`
}

// @BasePath /api/v1
// @Summary List users
// @Description List all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} ListUsersResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func ListUsersController(ctx *gin.Context) {
	users := []model.User{}

	// Find users
	if err := db.Find(&users).Error; err != nil {
		log.Println(err)
		sendError(ctx, http.StatusInternalServerError, "failed to list users")
		return
	}

	response := struct {
		Data []model.User
	}{
		Data: users,
	}

	ctx.JSON(http.StatusOK, response)
}
