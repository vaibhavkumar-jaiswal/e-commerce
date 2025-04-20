// Package dtos provides Data Transfer Objects (DTOs) for the product module.
package dtos

import "github.com/google/uuid"

// ProductRequest defines expected data when creating a new product
type ProductRequest struct {
	Name              string    `json:"name" example:"iPhone 14" validate:"required,alpha,min=2,max=50"`
	Description       string    `json:"description" example:"Latest Apple smartphone" validate:"min=2,max=500"`
	Price             float64   `json:"price" example:"999.99" validate:"required"`
	ProductCategoryID uuid.UUID `json:"product_category_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479" validate:"required"`
	Stock             int       `json:"stock" example:"100" validate:"required"`
	ImageURL          string    `json:"imageurl" example:"https://example.com/image.jpg"`
}

// ProductQueryParams defines allowed query parameters for filtering products
type ProductQueryParams struct {
	Name              string    `json:"name" example:"iPhone 14"`
	Price             float64   `json:"price" example:"999.99"`
	ProductCategoryID uuid.UUID `json:"product_category_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479"`
}

// UpdateProductRequest defines expected data when updating a product
type UpdateProductRequest struct {
	Name              string    `json:"name" validate:"required,alpha,min=2,max=50"`
	Description       string    `json:"description" validate:"omitempty,min=2,max=500"`
	Price             float64   `json:"price" validate:"required"`
	ProductCategoryID uuid.UUID `json:"product_category_id" validate:"required"`
	Stock             int       `json:"stock" validate:"required"`
	ImageURL          string    `json:"imageurl" validate:"omitempty"`
}

// PatchProductRequest defines expected data when updating a product with specific fields
type PatchProductRequest struct {
	Name              *string    `json:"name" example:"iPhone 14 Pro" validate:"omitempty,alpha,min=2,max=50"`
	Description       *string    `json:"description" example:"Updated description" validate:"omitempty,min=2,max=500"`
	Price             *float64   `json:"price" example:"1099.99" validate:"omitempty"`
	ProductCategoryID *uuid.UUID `json:"product_category_id" example:"f47ac10b-58cc-4372-a567-0e02b2c3d479" validate:"omitempty"`
	Stock             *int       `json:"stock" validate:"omitempty" example:"50"`
	ImageURL          *string    `json:"imageurl" validate:"omitempty" example:"https://example.com/image.jpg"`
}

// AddProductSuccess represents a successful creation of product response.
type AddProductSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Product added successfully."`
} // @name AddProductSuccess

// UpdateProductSuccess represents a successful updation of product response.
type UpdateProductSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Product upadated successfully."`
} // @name UpdateProductSuccess

// DeleteProductSuccess represents a successful deletion of product response.
type DeleteProductSuccess struct {
	Success bool   `json:"success" example:"true"`
	Data    string `json:"data" example:"Product successfully deleted."`
} // @name DeleteProductSuccess
