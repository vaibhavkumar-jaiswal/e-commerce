// Package handler provides the HTTP handlers for product management operations.
package handler

import (
	_ "e-commerce/models" // for swagger documentation
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
//	@Summary		Get list of products
//	@Description	Retrieves a list of products based on optional filters such as name, price, product category ID.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//
//	@Security		BearerAuth
//
//	@Param			name				query		string						false	"Filter by product name (partial match)"
//	@Param			price				query		number						false	"Filter by product price"
//	@Param			product_category_id	query		string						false	"Filter by Product Category ID (UUID)"
//	@Success		200					{object}	models.ProductResponse		"List of products"
//
//	@Failure		400					{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401					{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500					{object}	shared.InternalServerError	"Unexpected server error"
//
//	@Router			/product [get]
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

// AddProduct godoc
//
//	@Summary		Add a new product
//	@Description	Creates a new product with the provided details.
//	@Tags			Product
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			product	body		dtos.ProductRequest			true	"Product data"
//
//	@Success		200		{object}	dtos.AddProductSuccess		"Product created successfully"
//
//	@Failure		400		{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401		{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500		{object}	shared.InternalServerError	"Unexpected server error"
//
//	@Router			/product [post]
func (handler *Handler) AddProduct(context *gin.Context) {

	request, validationErr := helper.BindAndValidate[dtos.ProductRequest](context)
	if validationErr != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, validationErr)
		return
	}

	message, err := handler.service.AddProduct(&request)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}

// GetProductByID godoc
//
//	@Summary		Get a product by ID
//	@Description	Retrieves a single product by its unique identifier (UUID).
//	@Tags			Product
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string						true	"Product ID (UUID)"
//	@Success		200			{object}	models.ProductResponse		"Product details"
//
//	@Failure		400			{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401			{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500			{object}	shared.InternalServerError	"Unexpected server error"
//
//	@Router			/product/{product_id} [get]
func (handler *Handler) GetProductByID(context *gin.Context) {

	productID := context.Param("product_id")
	if productID == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID format.")
		return
	}
	user, err := handler.service.GetProductByID(productID)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, user)
}

// UpdateProduct godoc
//
//	@Summary		Update a product by ID
//	@Description	Updates an existing product using its unique ID. All fields will be updated.
//	@Tags			Product
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string						true	"Product ID (UUID)"
//	@Param			product		body		dtos.ProductRequest			true	"Updated product data"
//	@Success		200			{object}	dtos.UpdateProductSuccess	"Product updated successfully"
//
//	@Failure		400			{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401			{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500			{object}	shared.InternalServerError	"Unexpected server error"
//
//	@Router			/product/{product_id} [put]
func (handler *Handler) UpdateProduct(context *gin.Context) {
	helper.HandleUpdateAndPatch(
		"product_id",
		context,
		handler.service.UpdateProduct,
	)
}

// PartialUpdateProduct godoc
//
//	@Summary		Partially update a product by ID
//	@Description	Partially updates an existing product using its unique ID. Only the provided fields will be updated.
//	@Tags			Product
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//	@Param			product_id	path		string						true	"Product ID (UUID)"
//	@Param			product		body		dtos.PatchProductRequest	true	"Fields to update in the product"
//	@Success		200			{object}	dtos.UpdateProductSuccess	"Product updated successfully"
//
//	@Failure		400			{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401			{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500			{object}	shared.InternalServerError	"Unexpected server error"
//
//	@Router			/product/{product_id} [patch]
func (handler *Handler) PartialUpdateProduct(context *gin.Context) {
	helper.HandleUpdateAndPatch(
		"product_id",
		context,
		handler.service.PartialUpdateProduct,
	)
}

// DeleteProduct godoc
//
//	@Summary		Delete a product by ID
//	@Description	Deletes a product using its unique identifier (UUID).
//	@Tags			Product
//	@Security		BearerAuth
//	@Accept			json
//	@Produce		json
//
//	@Security		BearerAuth
//	@example		"Authorization: Bearer <your-jwt-token>"
//
//	@Param			product_id	path		string						true	"Product ID (UUID)"
//	@Success		200			{object}	dtos.DeleteProductSuccess	"Product deleted successfully"
//
//	@Failure		400			{object}	shared.BadRequestError		"Invalid input or missing required fields"
//	@Failure		401			{object}	shared.UnauthorizedError	"Unauthorized access attempt"
//	@Failure		500			{object}	shared.InternalServerError	"Unexpected server error"
//
//	@Router			/product/{product_id} [delete]
func (handler *Handler) DeleteProduct(context *gin.Context) {
	productID := context.Param("product_id")
	if productID == "" {
		helper.ResponseWriter(context, http.StatusBadRequest, "Invalid ID.")
		return
	}

	message, err := handler.service.DeleteProduct(productID)
	if err != nil {
		helper.ResponseWriter(context, http.StatusBadRequest, err.Error())
		return
	}

	helper.ResponseWriter(context, http.StatusOK, message)
}
