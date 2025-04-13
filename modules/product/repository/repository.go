// Package repository provides the repository layer for product management operations.
package repository

import (
	"e-commerce/database/connections"

	"e-commerce/base"
	"e-commerce/models"
)

// Repo defines a concrete implementation of user-specific repository
// using the generic BaseRepository from the shared layer.
type Repo struct {
	base base.Repository[models.Product]
}

// NewUserRepository creates a new instance of the User repository.
// Returns:
// - *Repo: Pointer to a new Repo with injected DB and Redis client.
func NewUserRepository() *Repo {
	return &Repo{
		base: *base.NewRepository[models.Product](connections.GetDB(), connections.GetRedisClient()),
	}
}
