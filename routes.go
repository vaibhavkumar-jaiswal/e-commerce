package main

import (
	product "e-commerce/modules/product/route"
	user "e-commerce/modules/user/route"

	"github.com/gin-gonic/gin"
)

// registerRoute registers all routes for the application
func registerRoute(router *gin.Engine) {
	user.UserManagementRoutes(router)
	product.ProductRoutes(router)
}
