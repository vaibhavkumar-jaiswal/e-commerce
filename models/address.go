// Package models contains the Address related models
package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

// Address represents a physical address associated with a user.
type Address struct {
	base.Model    `swaggerignore:"true"`
	AddressID     uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"address_id"`
	UserID        uuid.UUID   `gorm:"not null"`
	Street        string      `gorm:"size:255;not null"`
	Street2       string      `gorm:"size:255"`
	City          string      `gorm:"size:100;not null"`
	State         string      `gorm:"size:100;not null"`
	PostalCode    string      `gorm:"size:20;not null"`
	Country       string      `gorm:"size:100;not null"`
	AddressTypeID uuid.UUID   `gorm:"not null" json:"address_type_id"`
	AddressType   AddressType `gorm:"foreignKey:AddressTypeID;references:AddressTypeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	IsPrimary     bool        `gorm:"default:false"`
}

// TableName sets custom table name for Address model
func (Address) TableName() string {
	return "addresses"
}
