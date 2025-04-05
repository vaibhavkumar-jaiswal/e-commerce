package helper

import (
	cryptRand "crypto/rand"
	"e-commerce/database/connections"
	"e-commerce/shared/models"
	"e-commerce/utils/constants"
	"encoding/json"
	"fmt"
	"math/big"
	mathRand "math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var ExpiryTime int
var OtpExpTime int
var passwordLength int
var otpLength int
var redisClient *redis.Client

func InitiateHelper(config models.ConfigData) {
	ExpiryTime = config.SessionTimeOutmin
	OtpExpTime = config.OtpExpMin
	passwordLength = config.PasswordLength
	otpLength = config.OTPLength
	redisClient = connections.GetRedisClient()
}

// json related functions
func StructToJson(data any) string {
	// Marshal the embedding struct into JSON
	v, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return string(v)
}

func JsonToStruct[T any](data string) (T, error) {
	var v T
	// Marshal the embedding JSON into struct
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		return v, err
	}
	return v, nil
}

func CalculateOffset(page, limit string) int {
	pageInt := StringToInt(page)
	limitInt := StringToInt(limit)

	if pageInt <= 0 || limitInt <= 0 {
		return 0
	}

	offset := (pageInt - 1) * limitInt
	return offset
}

func StringToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Error while converting String to Int: %#v", err)
		return 0
	}
	return result
}

func IntToString(num int) string {
	return strconv.Itoa(num)
}

func FloatToInt(num float64) int {
	return 0
}

func CustomRecovery(context *gin.Context) {
	if r := recover(); r != nil {
		fmt.Printf("\n-------------------exception-------------------: \n%#v", r)
		context.JSON(500, gin.H{"data": "Internal Server Error"})
		return
	}
}

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

func BuildQuery(query *gorm.DB, filterStruct interface{}) *gorm.DB {

	// Use reflection to iterate over the fields of the filter struct
	v := reflect.ValueOf(filterStruct)
	t := reflect.TypeOf(filterStruct)

	// Handle pointer structs
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}

	// Iterate over the fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		valueField := v.Field(i)

		if !valueField.CanInterface() {
			continue
		}

		// Extract field value
		var value interface{}
		if valueField.Kind() == reflect.Ptr {
			// If the field is a pointer, check for nil and dereference
			if !valueField.IsNil() {
				value = valueField.Elem().Interface()
			} else {
				continue // Skip nil pointers
			}
		} else {
			value = valueField.Interface()
		}

		// Check if the field has a "query" tag and the value is not empty/zero
		tag := field.Tag.Get("form")
		qTag := field.Tag.Get("query")

		if (tag != "" && !isZero(value)) || (reflect.TypeOf(value).Kind() == reflect.Bool) {
			switch qTag {
			case "LIKE":
				query = query.Where(fmt.Sprintf("%s LIKE ?", tag), fmt.Sprintf("%%%s%%", value))
			case "ILIKE":
				query = query.Where(fmt.Sprintf("%s ILIKE ?", tag), fmt.Sprintf("%%%s%%", value))
			case "EQUAL":
				query = query.Where(fmt.Sprintf("%s = ?", tag), value)
			default:
				query = query.Where(fmt.Sprintf("%s = ?", tag), value)
			}
		}
	}

	return query
}

// Helper function to check if a value is the zero value for its type
func isZero(value interface{}) bool {
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
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
