package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// ProductCategory represents a product category in the e-commerce system.
type ProductCategory struct {
	base.Model        `swaggerignore:"true"`
	ProductCategoryID uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"product_category_id"`
	ParentID          *uuid.UUID        `gorm:"type:uuid;index" json:"parent_id"`
	Parent            *ProductCategory  `gorm:"foreignKey:ParentID" json:"-"`
	Children          []ProductCategory `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	Name              string            `gorm:"size:100;uniqueIndex;not null"`
	Code              string            `gorm:"not null;unique" json:"code"`
	Description       string            `gorm:"type:text"`
} // @name ProductCategory

// TableName sets custom table name for ProductCategory model
func (ProductCategory) TableName() string {
	return "product_categories"
}
