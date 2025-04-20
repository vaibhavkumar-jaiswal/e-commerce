// Package constants provides a collection of constant values used for configuration and
// environment variables throughout the application.
package constants

// Environment constants
const (
	LocalEnv = "local"
	DevEnv   = "dev"
	TestEnv  = "test"
)

// Environment variable keys
const (
	DbUser             = "DB_USER"
	DbPassword         = "DB_PASSWORD"
	AppEnv             = "APP_ENV"
	AppPort            = "APP_PORT"
	RedisHost          = "REDIS_HOST"
	RedisPort          = "REDIS_PORT"
	LogLevel           = "LOG_LEVEL"
	SMTPUser           = "EMAIL_USER"
	SMTPPassword       = "EMAIL_PASSWORD"
	EmailFrom          = "EMAIL_FROM"
	CompanyName        = "COMPANY_NAME"
	BasePath           = "BASE_PATH"
	PathLevelRateLimit = "PATH_LEVEL_RATE_LIMIT"
)
