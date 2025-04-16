// Package dtos provides Data Transfer Objects (DTOs) for the product module.
package dtos

import "github.com/google/uuid"

// ProductQueryParams defines allowed query parameters for filtering products
type ProductQueryParams struct {
	Name              string    `json:"name"`
	Price             float64   `json:"price"`
	ProductCategoryID uuid.UUID `json:"product_category_id"`
	SortBy            string    `json:"sort_by"`
}
