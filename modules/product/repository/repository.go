// Package repository provides the repository layer for product management operations.
package repository

import (
	"e-commerce/database/connections"

	"e-commerce/base"
	"e-commerce/models"

	"gorm.io/gorm"
)

// ProductRepo defines a concrete implementation of product-specific repository
// using the generic BaseRepository from the shared layer.
type ProductRepo struct {
	base base.Repository[models.Product]
}

// NewUserRepository creates a new instance of the Product repository.
// Returns:
// - *Repo: Pointer to a new Repo with injected DB and Redis client.
func NewProductRepository() *ProductRepo {
	return &ProductRepo{
		base: *base.NewRepository[models.Product](
			connections.GetDB(),
			connections.GetRedisClient(),
		),
	}
}

// GetFilter returns a GORM query builder for the Product model.
// This is useful for adding dynamic filters in services/controllers.
// Returns:
// - *gorm.DB: GORM DB model scoped to Product.
func (productRepo ProductRepo) GetFilter() *gorm.DB {
	return productRepo.base.DB.Model(&models.Product{})
}

// Get retrieves a product by their ID.
// Parameters:
// - id (uint): The product ID.
// Returns:
// - *models.Product: Pointer to the retrieved product, or nil if not found.
// - error: Error if any occurred during the query.
func (productRepo ProductRepo) Get(id uint) (*models.Product, error) {
	return productRepo.base.Get(id)
}

// FindAll retrieves all products from the database.
// It accepts filters, orderBy, limit, and offset parameters for pagination.
func (productRepo ProductRepo) FindAll(filters *gorm.DB, orderBy string, limit int, offset int) ([]models.Product, int64, error) {
	return productRepo.base.FindAll(filters, orderBy, limit, offset)
}

func (productRepo ProductRepo) Create(product *models.Product) error {
	return productRepo.base.Create(product)
}

func (productRepo ProductRepo) Update(product *models.Product) error {
	return productRepo.base.Update(product)
}

func (productRepo ProductRepo) PartialUpdate(record map[string]any, condition string, args ...any) error {
	return productRepo.base.PartialUpdate(record, condition, args...)
}

func (productRepo ProductRepo) Delete(product *models.Product, isSoftDelete bool) error {
	return productRepo.base.Delete(product, isSoftDelete)
}

// -------------------------------------Product Category-------------------------------------

// ProductCategoryRepo defines a concrete implementation of product-specific repository
// using the generic BaseRepository from the shared layer.
type ProductCategoryRepo struct {
	base base.Repository[models.ProductCategory]
}

// NewUserRepository creates a new instance of the Product repository.
// Returns:
// - *Repo: Pointer to a new Repo with injected DB and Redis client.
func NewProductCategoryRepository() *ProductCategoryRepo {
	return &ProductCategoryRepo{
		base: *base.NewRepository[models.ProductCategory](
			connections.GetDB(),
			connections.GetRedisClient(),
		),
	}
}

// GetFilter returns a GORM query builder for the ProductCategory model.
// This is useful for adding dynamic filters in services/controllers.
// Returns:
// - *gorm.DB: GORM DB model scoped to ProductCategory.
func (productCategoryRepo ProductCategoryRepo) GetFilter() *gorm.DB {
	return productCategoryRepo.base.DB.Model(&models.ProductCategory{})
}

// FindAll retrieves all product categories from the database.
// It accepts filters, orderBy, limit, and offset parameters for pagination.
func (productCategoryRepo *ProductCategoryRepo) FindAll(filters *gorm.DB, orderBy string, limit int, offset int) ([]models.ProductCategory, int64, error) {
	return productCategoryRepo.base.FindAll(filters, orderBy, limit, offset)
}
