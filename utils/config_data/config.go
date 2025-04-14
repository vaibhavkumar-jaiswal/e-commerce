// Package configdata provides a handler to preload static data into the database from JSON files.
package configdata

import (
	"e-commerce/database/connections"
	"e-commerce/models"
	_ "e-commerce/shared" // for swagger documentation
	"e-commerce/utils/helper"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gin-gonic/gin"
)

// PreLoadDataHandler godoc
//
//	@Summary		Preload Static Data
//	@Description	Loads static reference data into the database from JSON files.
//	@Tags			Admin
//	@Produce		json
//	@Success		200	{object}	shared.LoadDataSuccess		"Data loaded successfully from JSON files"
//	@Failure		500	{object}	shared.InternalServerError	"Error occurred while loading data"
//	@Router			/load-data [get]
func PreLoadDataHandler(context *gin.Context) {

	fileNames := []string{
		"product_category",
		"address_type",
		"user_role",
		"product",
		"user",
	}

	fileNameModels := map[string]any{
		"product_category": models.ProductCategory{},
		"address_type":     models.AddressType{},
		"product":          models.Product{},
		"user_role":        models.Role{},
		"user":             models.User{},
	}

	db := connections.GetDB()

	_ = db.Unscoped().Where("1 = 1").Delete(&models.UserPassword{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.Address{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.Product{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.ProductCategory{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.AddressType{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.User{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.Role{}).Error

	for _, fileName := range fileNames {
		// Get the current working directory.
		currentDir, err := os.Getwd()
		if err != nil {
			helper.ResponseWriter(context, http.StatusBadRequest, "Cannot get current working directory.")
			return
		}

		// Build the file location.
		fileLocation := filepath.Join(currentDir, "utils", "config_data", fileName+".json")

		// Open the JSON file.
		file, err := os.Open(fileLocation) // #nosec G304
		if err != nil {
			helper.ResponseWriter(
				context,
				http.StatusBadRequest,
				fmt.Sprintf("Failed to locate/open JSON file (%s)", fileLocation),
			)
			return
		}
		defer func() {
			err := file.Close()
			if err != nil {
				helper.ResponseWriter(
					context,
					http.StatusBadRequest,
					fmt.Sprintf("Failed to close JSON file (%s)", fileLocation),
				)
				return
			}
		}()

		// Use reflection to create a new slice of the model type.
		model := fileNameModels[fileName]
		modelType := reflect.TypeOf(model)
		modelSlice := reflect.MakeSlice(reflect.SliceOf(modelType), 0, 0)
		modelSlicePtr := reflect.New(modelSlice.Type())

		// Decode the JSON data into the slice.
		modelSlicePtr.Elem().Set(modelSlice)
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(modelSlicePtr.Interface()); err != nil {
			helper.ResponseWriter(
				context,
				http.StatusBadRequest,
				fmt.Sprintf("Failed to decode JSON data for %s: %v", fileName, err),
			)
			return
		}

		// Convert the reflected slice to a concrete slice for GORM bulk insertion.
		dataToInsert := modelSlicePtr.Elem().Interface()

		// Perform bulk insertion.
		result := db.Create(dataToInsert)

		// Check for errors during insertion.
		if result.Error != nil {
			fmt.Printf("Error inserting data for %s: %v\n", fileName, result.Error)
			helper.ResponseWriter(context, http.StatusBadRequest, fmt.Sprintf("Error inserting data for %s", fileName))
			return
		}

		fmt.Printf("Bulk data inserted successfully for %s\n", fileName)
	}

	helper.ResponseWriter(context, http.StatusOK, "Data inserted successfully!")
}
