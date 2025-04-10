package helper

import (
	"e-commerce/database/connections"
	"e-commerce/shared/models"
	"e-commerce/utils/constants"

	cryptRand "crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
)

var ExpiryTime int
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

// it converts the embedding struct into JSON format
// and returns the JSON string
func StructToJson(data any) string {
	v, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return string(v)
}

// it converts the JSON string into the embedding struct
// and returns the struct
func JsonToStruct[T any](data string) (T, error) {
	var v T
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		return v, err
	}
	return v, nil
}

// it calculates the offset for pagination
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

// it converts the string to int
// and returns the int value
func StringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error while converting String to Int: %#v", err)
		return 0
	}
	return result
}

// it converts the int to string
// and returns the string value
func IntToString(num int) string {
	return strconv.Itoa(num)
}

// it converts the float to int
// and returns the int value
// Note: This function currently returns 0 as the implementation is not provided.
// You can implement the conversion logic as per your requirement.
func FloatToInt(num float64) int {
	return 0
}

// it is used to recover from panics in the application.
// It logs the panic message and returns a JSON response with a 500 status code.
func CustomRecovery(context *gin.Context) {
	if r := recover(); r != nil {
		fmt.Printf("\n-------------------exception-------------------: \n%#v", r)
		context.JSON(500, gin.H{"data": "Internal Server Error"})
		return
	}
}

// it is used to create a JWT token with claims.
// It takes the data as input and returns the JWT token and a boolean indicating success or failure.
func CreateJwtWithClaims(data any) (string, bool) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[constants.USER_JWT_CLAIM_KEY] = data

	// Set token expiration time (e.g., 1 hour from now)
	expirationTime := time.Now().Add(time.Duration(ExpiryTime) * time.Minute)
	claims["exp"] = expirationTime.Unix()

	jwtToken, err := token.SignedString([]byte(os.Getenv(constants.SECRETE_KEY)))
	if err != nil {
		return "Failed to generate auth token", false
	}

	return jwtToken, true
}

// it is used to write the response to the client.
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

	// Create a new random generator with its own seed
	r := mathRand.New(mathRand.NewSource(time.Now().UnixNano()))

	// Create the password
	password := make([]byte, passwordLength)
	for i := range password {
		password[i] = charset[r.Intn(len(charset))]
	}
	return string(password)
}

// get email verification email format (subject, emailbody)
func GetEmailVerificationFormat(emailToName string, otp string, isHtml bool) (string, string) {
	companyName := os.Getenv(constants.COMPANY_NAME)
	subject := fmt.Sprintf(constants.OTP_VERIFICATION_EMAIL_SUBJECT, companyName)
	if isHtml {
		currentYear := time.Now().Year()
		return subject, fmt.Sprintf(constants.OTP_VERIFICATION_EMAIL_FORMAT_HTML, companyName, emailToName, otp, currentYear)
	} else {
		return subject, fmt.Sprintf(constants.OTP_VERIFICATION_EMAIL_FORMAT_TXT, emailToName, companyName, otp, companyName, companyName, companyName)
	}
}

// get login credential email format (subject, emailbody)
func GetCredentialEmailFormat(emailToName string, userID string, password string, isHtml bool) (string, string) {
	companyName := os.Getenv(constants.COMPANY_NAME)
	subject := fmt.Sprintf(constants.SHARE_CREDENTIAL_EMAIL_SUBJECT, companyName)
	if isHtml {
		currentYear := time.Now().Year()
		return subject, fmt.Sprintf(constants.SHARE_CREDENTIAL_EMAIL_FORMAT_HTML, companyName, emailToName, userID, password, companyName, currentYear)
	} else {
		return subject, fmt.Sprintf(constants.SHARE_CREDENTIAL_EMAIL_FORMAT_TXT, emailToName, companyName, userID, password, companyName, companyName)
	}

}

// GenerateOTP generates a numeric OTP of the specified length
func GenerateSecureOTP() string {
	const digits = "0123456789"
	otp := make([]byte, otpLength)

	for i := range otp {
		// Generate a secure random index
		num, err := cryptRand.Int(cryptRand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			fmt.Println("Error generating random number:", err)
			return ""
		}
		// Map the random index to a digit
		otp[i] = digits[num.Int64()]
	}

	return string(otp)
}
