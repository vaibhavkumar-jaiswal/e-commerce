// Package repository provides the repository layer for product management operations.
package repository

import (
	"e-commerce/database/connections"

	"e-commerce/base"
	"e-commerce/models"

	"gorm.io/gorm"
)

// Repo defines a concrete implementation of product-specific repository
// using the generic BaseRepository from the shared layer.
type Repo struct {
	productRepo         base.Repository[models.Product]
	productCategoryRepo base.Repository[models.ProductCategory]
}

// NewUserRepository creates a new instance of the Product repository.
// Returns:
// - *Repo: Pointer to a new Repo with injected DB and Redis client.
func NewUserRepository() *Repo {
	return &Repo{
		productRepo:         *base.NewRepository[models.Product](connections.GetDB(), connections.GetRedisClient()),
		productCategoryRepo: *base.NewRepository[models.ProductCategory](connections.GetDB(), connections.GetRedisClient()),
	}
}

// GetFilter returns a GORM query builder for the User model.
// This is useful for adding dynamic filters in services/controllers.
// Returns:
// - *gorm.DB: GORM DB model scoped to User.
func (repo Repo) GetFilter() *gorm.DB {
	return repo.productRepo.DB.Model(&models.Product{})
}

// FindAll retrieves all products from the database.
// It accepts filters, orderBy, limit, and offset parameters for pagination.
func (repo *Repo) FindAll(filters *gorm.DB, orderBy string, limit int, offset int) ([]models.Product, int64, error) {
	return repo.productRepo.FindAll(filters, orderBy, limit, offset)
}
