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

	modelMap := map[string]any{
		"Module":          &models.Module{},
		"Permission":      &models.Permission{},
		"Role":            &models.Role{},
		"User":            &models.User{},
		"ProductCategory": &models.ProductCategory{},
		"Product":         &models.Product{},
		"AddressType":     &models.AddressType{},
		"Address":         &models.Address{},
		"UserPassword":    &models.UserPassword{},
	}

	modelSequence := []string{
		"Module",
		"Permission",
		"Role",
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
