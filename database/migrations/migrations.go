// Package migrations handles database migrations for the application models.
package migrations

import (
	"e-commerce/database/connections"
	"e-commerce/models"
	"fmt"
)

// RunMigrations runs database migrations for the defined models. It uses GORM's AutoMigrate method
// to ensure the database schema is up to date with the model definitions.
func RunMigrations() error {
	db := connections.GetDB()

	modelsToMigrate := []any{
		&models.Role{},
		&models.User{},
		&models.UserPassword{},
		&models.AddressType{},
		&models.Address{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("migration failed for %T: %w", model, err)
		}
	}
	return nil
}
