// Package route provides the routing for the product module
package route

import (
	"e-commerce/middleware/auth"

	"github.com/gin-gonic/gin"
)

// ProductRoutes sets up the routes for product related operations.
func ProductRoutes(router *gin.Engine) {
	// handler := handler.NewUserHandler()

	router.GET(auth.PublicRoute("/product"), func(context *gin.Context) {
		context.JSON(200, "Product module is working")
	})

}
