package models

import (
	"e-commerce/base"

	"github.com/google/uuid"
)

type Address struct {
	base.BaseModel `swaggerignore:"true"`
	AddressID      uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"address_id"`
	UserID         uuid.UUID   `gorm:"not null"`
	Street         string      `gorm:"size:255;not null"`
	Street2        string      `gorm:"size:255"`
	City           string      `gorm:"size:100;not null"`
	State          string      `gorm:"size:100;not null"`
	PostalCode     string      `gorm:"size:20;not null"`
	Country        string      `gorm:"size:100;not null"`
	AddressTypeID  uuid.UUID   `gorm:"not null"`
	AddressType    AddressType `gorm:"foreignKey:AddressTypeID;references:AddressTypeID"`
	IsPrimary      bool        `gorm:"default:false"`
}

type AddressType struct {
	base.BaseModel `swaggerignore:"true"`
	AddressTypeID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"address_type_id"`
	Name           string    `gorm:"not null" json:"name"`
	Code           string    `gorm:"not null;unique" json:"code"`
	Description    string    `gorm:"size:255" json:"description"`
}
