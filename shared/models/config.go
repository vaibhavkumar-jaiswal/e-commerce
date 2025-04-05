package models

import (
	"fmt"
)

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
	SmtpServer        SmtpServer   `json:"smtp_server"`
	AllowedOrigins    []string     `json:"allowed_origins"`
	AllowedMethods    []string     `json:"allowed_methods"`
}

func (dbConnection *DBConnection) GetDBConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbConnection.Host, dbConnection.User, dbConnection.DBName, dbConnection.Password,
	)
}

type DBConnection struct {
	Host             string `json:"host"`
	User             string `json:"user"`
	DBName           string `json:"db_name"`
	SslMode          string `json:"ssl_mode"`
	Password         string `json:"password"`
	ConnectionString string `json:"connection_string"`
}

type Server struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type Logger struct {
	Request request `json:"request"`
	System  system  `json:"system"`
}

type SmtpServer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type request struct {
	LogDir         string `json:"logdir"`
	FilenamePrefix string `json:"filename_prefix"`
}

type system struct {
	LogDir         string `json:"logdir"`
	FilenamePrefix string `json:"filename_prefix"`
}

type RateLimit struct {
	MaxRequest int `json:"max_requests"`
	Duration   int `json:"duration_in_minute"`
}

type RedisConn struct {
	Address string `json:"address"`
	DB      int    `json:"db"`
}
