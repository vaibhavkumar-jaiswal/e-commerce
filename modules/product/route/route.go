// Package route provides the routing for the product module
package route

import (
	productHandler "e-commerce/modules/product/handler"

	"github.com/gin-gonic/gin"
)

// ProductRoutes sets up the routes for product related operations.
func ProductRoutes(router *gin.RouterGroup) {
	handler := productHandler.NewUserHandler()

	{
		router.GET("/product", handler.GetProducts)
	}

}
