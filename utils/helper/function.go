package helper

import (
	"e-commerce/database/connections"
	"e-commerce/shared"

	cryptRand "crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// ExpiryTime - jwt token expiry time in minutes
var ExpiryTime int

// OtpExpTime - OTP expiry time in minutes
var OtpExpTime int

var passwordLength int
var otpLength int
var redisClient *redis.Client

// InitiateHelper initializes the helper package with configuration data
// and sets up the Redis client.
func InitiateHelper(config shared.ConfigData) {
	ExpiryTime = config.SessionTimeOutmin
	OtpExpTime = config.OtpExpMin
	passwordLength = config.PasswordLength
	otpLength = config.OTPLength
	redisClient = connections.GetRedisClient()
}

// StructToJSON converts the embedding struct into JSON format
// and returns the JSON string
func StructToJSON(data any) string {
	v, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return string(v)
}

// JSONToStruct converts the JSON string into the embedding struct
// and returns the struct
func JSONToStruct[T any](data string) (T, error) {
	var v T
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		return v, err
	}
	return v, nil
}

// CalculateOffset calculates the offset for pagination
// based on the page number and limit provided.
func CalculateOffset(page, limit string) int {
	pageInt := StringToInt(page)
	limitInt := StringToInt(limit)

	if pageInt <= 0 || limitInt <= 0 {
		return 0
	}

	offset := (pageInt - 1) * limitInt
	return offset
}

// StringToInt converts the string to int
// and returns the int value
func StringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error while converting String to Int: %#v", err)
		return 0
	}
	return result
}

// IntToString converts the int to string
// and returns the string value
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// GeneratePassword generates a random password of the specified length.
// It includes uppercase letters, lowercase letters, digits, and special characters.
func GeneratePassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%^&*()-_=+[]{}<>?/"

	password := make([]byte, passwordLength)

	for i := range password {
		num, err := cryptRand.Int(cryptRand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			fmt.Println("Error generating random number:", err)
			return ""
		}

		password[i] = charset[num.Int64()]
	}

	return string(password)
}

// GenerateSecureOTP generates a numeric OTP of the specified length
func GenerateSecureOTP() string {
	const digits = "0123456789"
	otp := make([]byte, otpLength)

	for i := range otp {
		num, err := cryptRand.Int(cryptRand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			fmt.Println("Error generating random number:", err)
			return ""
		}

		otp[i] = digits[num.Int64()]
	}

	return string(otp)
}
