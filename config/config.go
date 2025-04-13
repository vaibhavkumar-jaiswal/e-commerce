// Package config contains the logic for loading and parsing the application's configuration.
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"e-commerce/models"
	"e-commerce/utils/constants"
)

// LoadConfig loads configuration settings from a JSON file based on the current environment.
// It returns a pointer to the ConfigData struct and any error encountered during the process.
func LoadConfig() (*models.ConfigData, error) {

	fileName := getConfigFileName()
	if fileName == "" {
		return nil, fmt.Errorf("please set a valid env & config file to get the config data")
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fileLocation := filepath.Join(currentDir, "config", fileName)

	file, err := os.Open(fileLocation) // #nosec G304
	if err != nil {
		return nil, fmt.Errorf("failed to locate/open JSON file (%s)", fileLocation)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	var configData models.ConfigData

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configData); err != nil {
		return nil, fmt.Errorf("failed to decode JSON data")
	}

	configData.DBConnection.User = os.Getenv(constants.DbUser)
	configData.DBConnection.Password = os.Getenv(constants.DbPassword)
	configData.SMTPServer.UserName = os.Getenv(constants.SMTPUser)
	configData.SMTPServer.Password = os.Getenv(constants.SMTPPassword)

	if os.Getenv(constants.AppEnv) == constants.LocalEnv {
		configData.SMTPServer.Port = 1025
		configData.SMTPServer.Host = "localhost"
		configData.SMTPServer.UserName = ""
		configData.SMTPServer.Password = ""
	}

	if os.Getenv(constants.AppEnv) == constants.LocalEnv {
		configData.SMTPServer.Port = 1025
		configData.SMTPServer.Host = "localhost"
		configData.SMTPServer.UserName = ""
		configData.SMTPServer.Password = ""
	}

	return &configData, nil
}

// getConfigFileName determines the appropriate configuration file based on the environment.
// It defaults to "local.json" if the environment variable is not set.
func getConfigFileName() string {
	if os.Getenv(constants.AppEnv) == "" {
		return "local.json"
	}
	return strings.ToLower(os.Getenv(constants.AppEnv)) + ".json"
}
