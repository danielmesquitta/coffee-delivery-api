package controller

import (
	"log"
	"net/http"

	"github.com/danielmesquitta/coffee-delivery-api/model"
	"github.com/gin-gonic/gin"
)

type ListUsersResponse struct {
	PaginatedResponse
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
func ListUsersController(c *gin.Context) {
	users := []model.User{}

	response := NewPaginatedResponse(c)

	// Find users
	if err := db.Scopes(Paginate(users, &response)).Find(&users).Error; err != nil {
		log.Println(err)
		sendError(c, http.StatusInternalServerError, "failed to list users")
		return
	}

	response.Data = users

	c.JSON(http.StatusOK, response)
}
