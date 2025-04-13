// Package models - product file defines the Product model and its methods.
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// Product represents a product in the e-commerce system.
type Product struct {
	base.Model  `swaggerignore:"true"`
	ProductID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"product_id"`
	Name        string    `gorm:"size:100;not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	CategoryID  uuid.UUID `gorm:"type:uuid;not null"`
	Category    Category  `gorm:"foreignKey:CategoryID"`
}

// Category represents a product category in the e-commerce system.
type Category struct {
	base.Model        `swaggerignore:"true"`
	ProductCategoryID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"product_category_id"`
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name              string    `gorm:"size:100;uniqueIndex;not null"`
}
