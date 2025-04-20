// Package route provides the routing for the product module
package route

import (
	productHandler "e-commerce/modules/product/handler"

	"github.com/gin-gonic/gin"
)

// ProductRoutes sets up the routes for product related operations.
func ProductRoutes(router *gin.RouterGroup) {
	handler := productHandler.NewUserHandler()

	product := router.Group("/product")
	{
		product.GET("", handler.GetProducts)

		product.GET("/:product_id", handler.GetProductByID)

		product.POST("", handler.AddProduct)

		product.PUT("/:product_id", handler.UpdateProduct)

		product.PATCH("/:product_id", handler.PartialUpdateProduct)

		product.DELETE("/:product_id", handler.DeleteProduct)
	}

}
