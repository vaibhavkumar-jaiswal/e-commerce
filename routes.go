package main

import (
	product "e-commerce/modules/product/route"
	user "e-commerce/modules/user/route"
	"e-commerce/utils/constants"
	"os"

	"github.com/gin-gonic/gin"
)

// registerRoute registers all routes for the application
func registerRoute(router *gin.Engine) {
	v1Router := router.Group(os.Getenv(constants.BasePath))
	user.UserManagementRoutes(v1Router)
	product.ProductRoutes(v1Router)
}
