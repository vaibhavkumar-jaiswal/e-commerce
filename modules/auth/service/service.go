// Package service provides user management operations, including login, email verification,
// user retrieval, registration, and update functionalities.
// It interacts with the repository layer to perform CRUD operations and other user-related queries.
// It also handles JWT token generation and email sending for user verification.
package service

import (
	"e-commerce/models"
	"e-commerce/modules/auth/dtos"
	"e-commerce/modules/auth/repository"
	"e-commerce/services"
	"e-commerce/utils/constants"
	"e-commerce/utils/helper"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"
)

// Service provides user management operations, including login, email verification,
// user retrieval, registration, and update functionalities.
type Service struct {
	userRepo *repository.UserRepo
}

// NewUserService creates and returns a new User Service instance by initializing the repository.
// Returns:
//
//	*Service: A pointer to a new Service instance with its repository initialized.
func NewUserService() *Service {
	return &Service{
		userRepo: repository.NewUserRepository(),
	}
}

// Login processes user authentication based on provided login credentials.
// It validates the username, password and then generates a JWT token upon successful authentication.
//
// Parameters:
//
//	data (models.Login): Login credentials containing UserName and Password.
//
// Returns:
//
//	any: Typically a models.LoginResponse on success.
//	error: An error if authentication fails or other issues occur.
func (service *Service) Login(data dtos.Login) (any, error) {
	join := "INNER JOIN user_passwords ON users.user_id = user_passwords.user_id"
	condition := "users.email = ? AND user_passwords.password = ? AND users.is_verified = true"

	relations := []string{"Role"}
	userList, err := service.userRepo.FindAllByConditionWithJoin(relations, join, condition, data.UserName, data.Password)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}
	if len(userList) < 1 {
		return nil, fmt.Errorf("invalid User credentials")
	}

	user := userList[0].ResponseObj()

	token, ok := helper.CreateJwtWithClaims(user)
	if !ok {
		return nil, fmt.Errorf("not able to create jwt token, please try again")
	}

	var response dtos.LoginResponse
	response.UserDetails = user
	response.AuthorizationToken = token
	response.Expiry = time.Now().Add(time.Duration(helper.ExpiryTime) * time.Minute)

	return response, nil
}

// Logout logs out a user by storing the provided token in a cache with an expiration.
// This prevents reuse of the same token after logout.
//
// Parameters:
//   - token: A string representing the user's JWT token to be invalidated.
//
// Returns:
//   - A success message string if logout is successful.
//   - An error if the token could not be cached for logout invalidation.
func (service *Service) Logout(token string) (string, error) {
	_, err := helper.SetCache(
		constants.LoggedOutKey+":"+token,
		constants.LoggedOutKey,
		time.Duration(helper.ExpiryTime)*time.Minute,
	)
	if err != nil {
		return "", fmt.Errorf("failed to log out user")
	}
	return "User logged out successfully", nil
}

// VerifyEmail verifies a user's email using an OTP (One-Time Password).
// It validates the user's status, compares the provided OTP with cached OTP,
// and, upon a successful match, updates the user's status to verified,
// clears the OTP cache, and sends a confirmation email.
//
// Parameters:
//
//	email (string): The email address to verify.
//	otp (string): The one-time password entered by the user.
//
// Returns:
//
//	any: A success message indicating the email is verified.
//	error: An error if verification fails or other issues occur.
func (service *Service) VerifyEmail(email, otp string) (any, error) {
	unverifiedUserQuery := "users.email = ?"

	user, err := service.userRepo.GetByCondition(unverifiedUserQuery, email)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("user with email (%s) is not registered", email)
	}

	if user.IsVerified {
		return nil, fmt.Errorf("user with email (%s) is already verified, you can proceed to login", email)
	}

	cachedOtp, err := helper.GetCache(email)
	if err != nil {
		return nil, fmt.Errorf("the OTP you entered is expired")
	}

	sentOtp := strings.ReplaceAll(otp, " ", "")

	if cachedOtp != sentOtp {
		return nil, fmt.Errorf("the OTP you entered is incorrect. Please check and try again")
	}

	record := map[string]any{
		"is_verified": true,
	}

	condition := "users.email = ? AND users.is_verified = false"

	err = service.userRepo.PartialUpdate(record, condition, email)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	relations := []string{"UserPassword", "Role"}
	join := "INNER JOIN user_passwords ON users.user_id = user_passwords.user_id"

	users, err := service.userRepo.FindAllByConditionWithJoin(
		relations,
		join,
		"users.email = ? AND users.is_verified = true",
		email,
	)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if len(users) < 1 {
		return nil, fmt.Errorf("user details not found")
	}

	err = helper.DeleteCache(email)
	if err != nil {
		return nil, fmt.Errorf("failed to delete cache")
	}

	isHTML := true
	subject, emailBody := helper.GetCredentialEmailFormat(
		users[0].FullName(),
		users[0].Email,
		users[0].UserPassword.Password,
		isHTML,
	)

	go func() {
		if err := services.SMTPServer.SendEmail(users[0].Email, subject, emailBody, isHTML); err != nil {
			fmt.Printf("failed to send email to %s: %v", users[0].Email, err)
		}
	}()

	return "Your email has been successfully verified! We've sent your login credentials to your " +
		"registered email address. Please check your inbox to proceed.", nil
}

// ResendVerificationCode handles the process to resend an OTP verification code
// to users who have not yet verified their email.
//
// Parameters:
//
//	email (string): The email address to which the OTP will be sent.
//
// Returns:
//
//	any: A string message confirming that the OTP has been sent.
//	error: An error if any step fails during processing.
func (service *Service) ResendVerificationCode(email string) (any, error) {
	condition := "users.email = ?"

	user, err := service.userRepo.GetByCondition(condition, email)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("the user with email (%s) has not registered", email)
	}

	if user.IsVerified {
		return nil, fmt.Errorf("user with email (%s) is already verified, you can proceed to login", email)
	}

	otp := helper.GenerateSecureOTP()
	isHTML := true
	subject, emailBody := helper.GetEmailVerificationFormat(user.FullName(), otp, isHTML)

	_, err = helper.SetCache(user.Email, otp, time.Duration(helper.OtpExpTime)*time.Minute)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := services.SMTPServer.SendEmail(user.Email, subject, emailBody, isHTML); err != nil {
			fmt.Printf("failed to send email to %s: %v", user.Email, err)
		}
	}()

	return "We have sent the OTP to your Email address.", nil
}

// AddUser creates a new user in the system and sends an email verification OTP.
// It populates the user data from the request, generates a password,
// and caches an OTP before sending it asynchronously.
//
// Parameters:
//
//	request (models.UserRequest): The user data for registration.
//
// Returns:
//
//	string: A success message instructing the user to verify their email.
//	error: An error if the creation or OTP email sending fails.
func (service *Service) AddUser(request dtos.UserRequest) (string, error) {
	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
		RoleID:    request.RoleID,
		UserPassword: models.UserPassword{
			Password: helper.GeneratePassword(),
		},
	}
	err := service.userRepo.Create(&user)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	otp := helper.GenerateSecureOTP()
	isHTML := true
	subject, emailBody := helper.GetEmailVerificationFormat(user.FullName(), otp, isHTML)

	_, err = helper.SetCache(user.Email, otp, time.Duration(helper.OtpExpTime)*time.Minute)
	if err != nil {
		return "", err
	}

	go func() {
		if err := services.SMTPServer.SendEmail(user.Email, subject, emailBody, isHTML); err != nil {
			fmt.Printf("failed to send email to %s: %v", user.Email, err)
		}
	}()

	return "We have sent the OTP to your Email address.", nil
}
