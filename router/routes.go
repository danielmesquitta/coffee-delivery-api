package router

import (
	"github.com/danielmesquitta/coffee-delivery-api/controller"
	"github.com/danielmesquitta/coffee-delivery-api/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/api/v1"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// Products
		v1.POST("/product", controller.CreateProductController)
		v1.PUT("/product", controller.UpdateProductController)
		v1.DELETE("/product", controller.DeleteProductController)
		v1.GET("/product", controller.ShowProductController)
		v1.GET("/products", controller.ListProductsController)
	}
	{
		// Users
		v1.POST("/user", controller.CreateUserController)
		v1.PUT("/user", controller.UpdateUserController)
		v1.DELETE("/user", controller.DeleteUserController)
		v1.GET("/user", controller.ShowUserController)
		v1.GET("/users", controller.ListUsersController)
	}
	{
		// Orders
		v1.POST("/order", controller.CreateOrderController)
		v1.GET("/order", controller.ShowOrderController)
		v1.GET("/orders", controller.ListOrdersController)
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
