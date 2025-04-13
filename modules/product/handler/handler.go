// Package handler provides the HTTP handlers for product management operations.
package handler

import "e-commerce/modules/product/service"

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
