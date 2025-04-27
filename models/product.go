// Package models - product file defines the Product model and its methods.
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// Product represents a product in the e-commerce system.
type Product struct {
	base.Model        `swaggerignore:"true"`
	ProductID         uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"product_id"`
	Name              string          `gorm:"size:100;not null" validate:"required" json:"name"`
	Description       string          `gorm:"type:text" json:"description"`
	Price             float64         `gorm:"not null" validate:"required" json:"price"`
	ProductCategoryID uuid.UUID       `gorm:"type:uuid;not null" validate:"required" json:"product_category_id"`
	Category          ProductCategory `gorm:"foreignKey:ProductCategoryID;references:ProductCategoryID" json:"-"`
	Stock             int             `gorm:"not null" validate:"required" json:"stock"`
	ImageURL          string          `gorm:"size:255" json:"image_url"`
}

// TableName sets custom table name for Product model
func (Product) TableName() string {
	return "products"
}

// ProductList defines a slice of Product
type ProductList []Product

// ProductResponse defines how product data is returned to clients
type ProductResponse struct {
	ProductID         uuid.UUID       `json:"product_id"`
	Name              string          `json:"name"`
	Description       string          `json:"description"`
	Price             float64         `json:"price"`
	ProductCategoryID uuid.UUID       `json:"product_category_id"`
	Category          ProductCategory `json:"-"`
	Stock             int             `json:"stock"`
	ImageURL          string          `json:"image_url"`
}

// ResponseObj formats a single product to ProductResponse
func (product Product) ResponseObj() ProductResponse {
	result := ProductResponse{
		ProductID:         product.ProductID,
		Name:              product.Name,
		Description:       product.Description,
		Price:             product.Price,
		ProductCategoryID: product.ProductCategoryID,
		Category:          product.Category,
		Stock:             product.Stock,
		ImageURL:          product.ImageURL,
	}

	return result
}

// ResponseList formats a list of products to a list of ProductResponse
func (productList ProductList) ResponseList() []ProductResponse {
	var result []ProductResponse
	for _, obj := range productList {
		result = append(result, obj.ResponseObj())
	}
	return result
}
