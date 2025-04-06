package configdata

import (
	"e-commerce/database/connections"
	"e-commerce/shared/models"
	"e-commerce/utils/helper"

	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Preload data from JSON files into the database
// @Summary      Load Data
// @Description  Preload data from JSON files into the database
// @Tags         Preload Data
// @Produce      json
// @Success      200  {object}  models.SuccessResponse[string]
// @Failure      500  {object}  models.ErrorResponse[string]
// @Router       /load-data [get]
func PreLoadDataHandler(context *gin.Context) {

	fileNames := []string{
		"address_type",
		"user_role",
		"user",
	}

	fileNameModels := map[string]any{
		"address_type": models.AddressType{},
		"user_role":    models.Role{},
		"user":         models.User{},
	}

	db := connections.GetDB()

	_ = db.Unscoped().Where("1 = 1").Delete(&models.UserPassword{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.Address{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.User{}).Error
	_ = db.Unscoped().Where("1 = 1").Delete(&models.AddressType{}).Error
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
		file, err := os.Open(fileLocation)
		if err != nil {
			helper.ResponseWriter(context, http.StatusBadRequest, fmt.Sprintf("Failed to locate/open JSON file (%s)", fileLocation))
			return
		}
		defer file.Close()

		// Use reflection to create a new slice of the model type.
		model := fileNameModels[fileName]
		modelType := reflect.TypeOf(model)
		modelSlice := reflect.MakeSlice(reflect.SliceOf(modelType), 0, 0)
		modelSlicePtr := reflect.New(modelSlice.Type())

		// Decode the JSON data into the slice.
		modelSlicePtr.Elem().Set(modelSlice)
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(modelSlicePtr.Interface()); err != nil {
			helper.ResponseWriter(context, http.StatusBadRequest, fmt.Sprintf("Failed to decode JSON data for %s: %v", fileName, err))
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

/*
func PreLoadDataHandler(context *gin.Context) {

	// fileNames := []string{
	// 	"address_type",
	// 	"user_role",
	// }

	fileNames := map[string]any{
		"address_type": []models.AddressType{},
		"user_role": []models.Role{},
	}

	for fileName, model := range fileNames {
		currentDir, err := os.Getwd()
		if err != nil {
			helper.ResponseWriter(context, http.StatusBadRequest, "Cannot get current working directory.")
			return
		}

		fileLocation := filepath.Join(currentDir, "config", fileName)

		file, err := os.Open(fileLocation)
		if err != nil {
			helper.ResponseWriter(context, http.StatusBadRequest, fmt.Sprintf("failed to locate/open JSON file (%s)", fileLocation))
			return
		}
		defer file.Close()

		// var users []model
		modelType__ := reflect.TypeOf(model)
		modelType := reflect.New(modelType__)

		var data []modelType
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&users); err != nil {
			panic("failed to decode JSON data")
		}
	}

	for _, user := range users {
		result := connections.GetDB().Create(&user)

		// Check for errors during the insertion.
		if result.Error != nil {
			fmt.Println("Error inserting data:", result.Error)
			context.JSON(http.StatusBadRequest, "Error while inserting data...!")
			context.Abort()
			return
		} else {
			fmt.Println("Data inserted successfully!")
		}
	}
	context.JSON(http.StatusOK, "Data inserted successfully...!")

}
*/
