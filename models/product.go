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
	Name              string          `gorm:"size:100;not null"`
	Description       string          `gorm:"type:text"`
	Price             float64         `gorm:"not null"`
	ProductCategoryID uuid.UUID       `gorm:"type:uuid;not null"`
	Category          ProductCategory `gorm:"foreignKey:ProductCategoryID;references:ProductCategoryID"`
	Stock             int             `gorm:"not null"`
	ImageURL          string          `gorm:"size:255"`
}

// TableName sets custom table name for Product model
func (Product) TableName() string {
	return "products"
}
