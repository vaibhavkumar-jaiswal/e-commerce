package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	// baseModel
	FirstName    string       `gorm:"not null" json:"first_name" validate:"required,alpha,min=2,max=50"`
	LastName     string       `gorm:"not null" json:"last_name" validate:"required,alpha,min=2,max=50"`
	Email        string       `gorm:"unique;index;not null" json:"email" validate:"required,email"`
	Phone        string       `gorm:"not null" json:"phone" validate:"required,numeric,len=10"`
	RoleID       uuid.UUID    `gorm:"not null" json:"role_id"`
	Role         Role         `gorm:"foreignKey:RoleID;references:ID"`
	IsVerified   bool         `gorm:"default:false" json:"is_verified"`
	UserPassword UserPassword `gorm:"foreignKey:UserID" json:"user_password"`
	// IsDeleted    bool         `gorm:"default:false" json:"is_deleted"`
}

type UserPassword struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsDeleted bool       `gorm:"default:false" json:"is_deleted"`
	DeletedAt *time.Time `json:"deleted_at"`
	// baseModel
	Password string    `json:"password"`
	UserID   uuid.UUID `json:"user_id"`
}

type UserList []User

type UserRequest struct {
	FirstName string    `json:"first_name" validate:"required,alpha,min=2,max=50"`
	LastName  string    `json:"last_name" validate:"required,alpha,min=2,max=50"`
	Email     string    `json:"email" validate:"required,email"`
	Phone     string    `json:"phone" validate:"required,numeric,len=10"`
	RoleID    uuid.UUID `json:"role_id" validate:"required"`
} //@name UserRequest

type UserQueryParams struct {
	FirstName  *string    `form:"first_name" query:"ILIKE"`
	LastName   *string    `form:"last_name" query:"ILIKE"`
	Email      *string    `form:"email"`
	Phone      *string    `form:"phone"`
	RoleID     *uuid.UUID `form:"role_id"`
	IsDeleted  *bool      `form:"is_deleted"`
	IsVerified bool       `form:"is_verified"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	FirstName string    `json:"first_name" example:"John"`
	FullName  string    `json:"full_name" example:"John Doe"`
	LastName  string    `json:"last_name" example:"Doe"`
	Email     string    `json:"email" example:"john.doe@gmail.com"`
	Phone     string    `json:"phone" example:"1234567890"`
	RoleID    uuid.UUID `json:"role_id" example:"97d699c0-24ff-48dc-b64a-c29353fa8865"`
	IsDeleted bool      `json:"is_deleted" example:"false"`
} //@name UserResponse

func (user User) ResponseObj() UserResponse {
	result := UserResponse{
		ID:        user.ID,
		FullName:  user.FullName(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		RoleID:    user.RoleID,
		IsDeleted: user.IsDeleted,
	}

	return result
}

func (userList UserList) ResponseList() []UserResponse {
	var result []UserResponse
	for _, obj := range userList {
		result = append(result, obj.ResponseObj())
	}
	return result
}

func (user *User) FullName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
