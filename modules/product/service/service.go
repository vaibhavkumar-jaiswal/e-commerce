// Package service provides the service layer for product management operations.
package service

import "e-commerce/modules/product/repository"

// Service defines the structure for the product service layer.
// It contains a repository instance to interact with the database.
type Service struct {
	repo *repository.Repo
}

// NewUserService creates and returns a new User Service instance by initializing the repository.
// Returns:
//
//	*Service: A pointer to a new Service instance with its repository initialized.
func NewUserService() *Service {
	repo := repository.NewUserRepository()
	return &Service{
		repo: repo,
	}
}
