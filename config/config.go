package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"e-commerce/shared/models"
	"e-commerce/utils/constants"
)

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

	file, err := os.Open(fileLocation)
	if err != nil {
		return nil, fmt.Errorf("failed to locate/open JSON file (%s)", fileLocation)
	}
	defer file.Close()

	var configData models.ConfigData

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configData); err != nil {
		return nil, fmt.Errorf("failed to decode JSON data")
	}

	configData.DBConnection.User = os.Getenv(constants.DB_USER)
	configData.DBConnection.Password = os.Getenv(constants.DB_PASSWORD)
	configData.SmtpServer.UserName = os.Getenv(constants.SMTP_USER)
	configData.SmtpServer.Password = os.Getenv(constants.SMTP_PASSWORD)

	if os.Getenv(constants.APP_ENV) == constants.LOCAL_ENV {
		configData.SmtpServer.Port = 1025
		configData.SmtpServer.Host = "localhost"
		configData.SmtpServer.UserName = ""
		configData.SmtpServer.Password = ""
	}

	return &configData, nil
}

func getConfigFileName() string {
	if os.Getenv(constants.APP_ENV) == "" {
		return "local.json"
	}
	return strings.ToLower(os.Getenv(constants.APP_ENV)) + ".json"
}
