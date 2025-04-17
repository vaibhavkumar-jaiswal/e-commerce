// Package dtos provides Data Transfer Objects (DTOs) for the product module.
package dtos

import "github.com/google/uuid"

// ProductRequest defines expected data when creating a new product
type ProductRequest struct {
	Name              string    `json:"name" validate:"required"`
	Description       string    `json:"description"`
	Price             float64   `json:"price" validate:"required"`
	ProductCategoryID uuid.UUID `json:"product_category_id" validate:"required"`
	Stock             int       `json:"stock" validate:"required"`
	ImageURL          string    `json:"imageurl"`
}

// ProductQueryParams defines allowed query parameters for filtering products
type ProductQueryParams struct {
	Name              string    `json:"name"`
	Price             float64   `json:"price"`
	ProductCategoryID uuid.UUID `json:"product_category_id"`
	SortBy            string    `json:"sort_by"`
}
