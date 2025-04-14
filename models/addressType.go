package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// AddressType defines categories for addresses (e.g., home, work, billing).
type AddressType struct {
	base.Model    `swaggerignore:"true"`
	AddressTypeID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;unique" json:"address_type_id"`
	Name          string    `gorm:"not null" json:"name"`
	Code          string    `gorm:"not null;unique" json:"code"`
	Description   string    `gorm:"size:255" json:"description"`
}

// TableName sets custom table name for AddressType model
func (AddressType) TableName() string {
	return "address_types"
}
