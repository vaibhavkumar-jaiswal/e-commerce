// Package migrations handles database migrations for the application models.
package migrations

import (
	"e-commerce/database/connections"
	"e-commerce/models"
	"fmt"
)

// RunMigrations runs GORM migrations for selected models.
// If no model names are provided, it migrates all.
func RunMigrations(dbModels ...string) error {
	db := connections.GetDB()

	if err := db.SetupJoinTable(&models.Role{}, "Permissions", &models.RolePermission{}); err != nil {
		return fmt.Errorf("failed to set up join table: %w", err)
	}

	modelMap := map[string]any{
		"Permission":      &models.Permission{},
		"Role":            &models.Role{},
		"RolePermission":  &models.RolePermission{},
		"User":            &models.User{},
		"ProductCategory": &models.ProductCategory{},
		"Product":         &models.Product{},
		"AddressType":     &models.AddressType{},
		"Address":         &models.Address{},
		"UserPassword":    &models.UserPassword{},
	}

	modelSequence := []string{
		"Permission",
		"Role",
		"RolePermission",
		"User",
		"ProductCategory",
		"Product",
		"AddressType",
		"Address",
		"UserPassword",
	}

	var modelsToMigrate []any

	if len(dbModels) == 0 {
		for _, key := range modelSequence {
			modelsToMigrate = append(modelsToMigrate, modelMap[key])
		}
	} else {
		for _, name := range dbModels {
			if model, ok := modelMap[name]; ok {
				modelsToMigrate = append(modelsToMigrate, model)
			} else {
				return fmt.Errorf("unknown model: %s", name)
			}
		}
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("migration failed for %T: %w", model, err)
		}
	}

	return nil
}
