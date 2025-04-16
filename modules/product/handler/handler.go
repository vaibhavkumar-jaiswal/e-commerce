// Package handler provides the HTTP handlers for product management operations.
package handler

import (
	"e-commerce/modules/product/dtos"
	"e-commerce/modules/product/service"
	_ "e-commerce/shared" // for swagger documentation
	"e-commerce/utils/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler is the struct that contains the product service
// and is responsible for handling product-related requests.
type Handler struct {
	service *service.Service
}

// NewUserHandler returns the product handler
func NewUserHandler() *Handler {
	service := service.NewUserService()
	return &Handler{
		service: service,
	}
}

// GetProducts godoc
//
//	@Summary		List products
//	@Description	Retrieve a list of products with optional filters like category, sorting,
//	@Description	and pagination
//
//	@Tags			Products
//	@Accept			json
//	@Produce		json
//	@Param			page				query		int		false	"Page number"				minimum(1)	default(1)
//	@Param			page_size			query		int		false	"Number of items per page"	minimum(1)	maximum(100)	default(10)
//	@Param			product_category_id	query		string	false	"Product Category filter"
//	@Param			sort_by				query		string	false	"Sort by field (e.g., price, name)"
//	@Success		200					{array}		models.ProductResponse
//	@Failure		400					{object}	shared.BadRequestError		"Invalid query parameters"
//	@Failure		500					{object}	shared.InternalServerError	"Internal server error"
//	@Router			/products [get]
func (handler *Handler) GetProducts(context *gin.Context) {

	queryParams := &dtos.ProductQueryParams{}

	if err := context.ShouldBindQuery(queryParams); err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	products, err := handler.service.GetProducts(queryParams)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, products)
}
