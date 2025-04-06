package models

import (
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	// baseModel
	UserID        uuid.UUID   `gorm:"not null"`
	Street        string      `gorm:"size:255;not null"`
	Street2       string      `gorm:"size:255"`
	City          string      `gorm:"size:100;not null"`
	State         string      `gorm:"size:100;not null"`
	PostalCode    string      `gorm:"size:20;not null"`
	Country       string      `gorm:"size:100;not null"`
	AddressTypeID uuid.UUID   `gorm:"not null"` // Foreign key to AddressType
	AddressType   AddressType `gorm:"foreignKey:AddressTypeID;references:ID"`
	IsPrimary     bool        `gorm:"default:false"`
	// IsDeleted     bool        `gorm:"default:false" json:"is_deleted"`
}

type AddressType struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	// baseModel
	Name        string `gorm:"not null" json:"name"`
	Code        string `gorm:"not null;unique" json:"code"`
	Description string `gorm:"size:255" json:"description"` // Optional description
	// IsDeleted   bool   `gorm:"default:false" json:"is_deleted"`
}
