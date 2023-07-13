package router

import (
	"github.com/danielmesquitta/coffee-delivery-api/controller"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"

	v1 := router.Group(basePath)
	{
		v1.POST("/product", controller.CreateProduct)
	}
}
