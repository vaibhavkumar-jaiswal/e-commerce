package helper

import (
	"e-commerce/database/connections"
	"e-commerce/models"
	"e-commerce/utils/constants"

	cryptRand "crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
func InitiateHelper(config models.ConfigData) {
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

// CreateJwtWithClaims is used to create a JWT token with claims.
// It takes the data as input and returns the JWT token and a boolean indicating success or failure.
func CreateJwtWithClaims(data any) (string, bool) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "Failed to create auth token", false
	}
	claims[constants.UserJwtClaimKey] = data

	// Set token expiration time (e.g., 1 hour from now)
	expirationTime := time.Now().Add(time.Duration(ExpiryTime) * time.Minute)
	claims["exp"] = expirationTime.Unix()

	jwtToken, err := token.SignedString([]byte(os.Getenv(constants.SecretKey)))
	if err != nil {
		return "Failed to generate auth token", false
	}

	return jwtToken, true
}

// ResponseWriter is used to write the response to the client.
// It takes the context, status code, and data as input.
func ResponseWriter[T any](cxt *gin.Context, status int, data T) {
	var response any
	if status >= http.StatusBadRequest {
		response = models.ErrorResponse[T]{
			Success: false,
			Error:   data,
		}
	} else {
		response = models.SuccessResponse[T]{
			Success: true,
			Data:    data,
		}
	}
	cxt.JSON(status, response)
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

// GetEmailVerificationFormat generates the email format for OTP verification and returns email subject and body.
func GetEmailVerificationFormat(emailToName string, otp string, isHTML bool) (string, string) {
	companyName := os.Getenv(constants.CompanyName)
	subject := fmt.Sprintf(constants.OtpVerificationEmailSubject, companyName)
	if isHTML {
		currentYear := time.Now().Year()
		return subject,
			fmt.Sprintf(
				constants.OtpVerificationEmailFormatHTML,
				companyName,
				emailToName,
				otp,
				currentYear,
			)
	}

	return subject,
		fmt.Sprintf(
			constants.OtpVerificationEmailFormatTxt,
			emailToName,
			companyName,
			otp,
			companyName,
			companyName,
			companyName,
		)
}

// GetCredentialEmailFormat generates the email format for sharing credentials and returns email subject and body.
func GetCredentialEmailFormat(emailToName string, userID string, password string, isHTML bool) (string, string) {
	companyName := os.Getenv(constants.CompanyName)
	subject := fmt.Sprintf(constants.ShareCredentialEmailSubject, companyName)
	if isHTML {
		currentYear := time.Now().Year()
		return subject,
			fmt.Sprintf(
				constants.ShareCredentialEmailFormatHTML,
				companyName,
				emailToName,
				userID,
				password,
				companyName,
				currentYear,
			)
	}

	return subject,
		fmt.Sprintf(
			constants.ShareCredentialEmailFormatTxt,
			emailToName,
			companyName,
			userID,
			password,
			companyName,
			companyName,
		)

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
