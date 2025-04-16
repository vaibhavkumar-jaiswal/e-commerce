package main

import (
	"e-commerce/middleware/auth"
	product "e-commerce/modules/product/route"
	user "e-commerce/modules/user/route"
	"e-commerce/utils/constants"
	"e-commerce/utils/loaddata"
	"os"

	"github.com/gin-gonic/gin"
)

// registerRoute registers all routes for the application
func registerRoute(router *gin.Engine) {
	v1Router := router.Group(os.Getenv(constants.BasePath))
	v1Router.GET(auth.PublicRoute("/load-data"), loaddata.PreLoadDataHandler)
	user.UserManagementRoutes(v1Router)
	product.ProductRoutes(v1Router)
}
