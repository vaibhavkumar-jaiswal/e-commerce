// Package shared contains configuration structs and associated methods
// that handle the application's environment settings and connections.
package shared

import (
	"fmt"
)

// ConfigData holds all configuration settings for the application,
// including database connection, server settings, Redis, logging, rate limiting, and more.
type ConfigData struct {
	DBConnection      DBConnection `json:"db_connection"`
	Server            Server       `json:"server"`
	RedisConnection   RedisConn    `json:"redis_connection"`
	Logger            Logger       `json:"logger"`
	RateLimit         RateLimit    `json:"rate_limit"`
	SessionTimeOutmin int          `json:"session_timeout_min"`
	OtpExpMin         int          `json:"otp_exp_min"`
	PasswordLength    int          `json:"password_length"`
	OTPLength         int          `json:"otp_length"`
	SMTPServer        SMTPServer   `json:"smtp_server"`
	AllowedOrigins    []string     `json:"allowed_origins"`
	AllowedMethods    []string     `json:"allowed_methods"`
}

// DBConnection represents the settings needed to establish a connection to a database.
type DBConnection struct {
	Host             string `json:"host"`
	User             string `json:"user"`
	DBName           string `json:"db_name"`
	SslMode          string `json:"ssl_mode"`
	Password         string `json:"password"`
	ConnectionString string `json:"connection_string"`
}

// GetDBConnectionString generates the connection string for the database
// based on the configuration fields in DBConnection.
func (dbConnection *DBConnection) GetDBConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbConnection.Host, dbConnection.User, dbConnection.DBName, dbConnection.Password,
	)
}

// Server holds the settings for the web server, including the address and port.
type Server struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

// Logger defines the logging configurations, split into request and system logs.
type Logger struct {
	Request request `json:"request"`
	System  system  `json:"system"`
}

// SMTPServer defines the configuration settings for an SMTP server.
type SMTPServer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

// request defines configuration settings for request-specific logging.
type request struct {
	LogDir         string `json:"logdir"`
	FilenamePrefix string `json:"filename_prefix"`
}

// system defines configuration settings for system-wide logging.
type system struct {
	LogDir         string `json:"logdir"`
	FilenamePrefix string `json:"filename_prefix"`
}

// RateLimit defines the configuration for rate limiting, such as max requests and duration.
type RateLimit struct {
	MaxRequest int `json:"max_requests"`
	Duration   int `json:"duration_in_minute"`
}

// RedisConn defines the configuration settings for connecting to a Redis instance.
type RedisConn struct {
	Address string `json:"address"`
	DB      int    `json:"db"`
}
