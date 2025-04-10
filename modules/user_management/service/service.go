package service

import (
	"e-commerce/modules/user_management/repo"
	"e-commerce/services"
	"e-commerce/shared/models"
	"e-commerce/utils/helper"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Service provides user management operations, including login, email verification,
// user retrieval, registration, and update functionalities.
type Service struct {
	repo *repo.Repo
}

// NewUserService creates and returns a new User Service instance by initializing the repository.
// Returns:
//
//	*Service: A pointer to a new Service instance with its repository initialized.
func NewUserService() *Service {
	repo := repo.NewUserRepository()
	return &Service{
		repo: repo,
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
func (service *Service) Login(data models.Login) (any, error) {
	join := "INNER JOIN user_passwords ON users.id = user_passwords.user_id"
	condition := "users.email = ? AND user_passwords.password = ? AND users.is_verified = true"

	relations := []string{"Role"}
	userList, err := service.repo.FindAllByConditionWithJoin(relations, join, condition, data.UserName, data.Password)
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

	var response models.LoginResponse
	response.UserDetails = user
	response.AuthorizationToken = token
	response.Expiry = time.Now().Add(time.Duration(helper.ExpiryTime) * time.Minute)

	return response, nil
}

// GetUserByID retrieves a user by their unique identifier (UUID).
//
// Parameters:
//
//	id (string): The UUID of the user in string format.
//
// Returns:
//
//	any: Typically a user response object if found.
//	error: An error if the user is not found or if an error occurred during retrieval.
func (service *Service) GetUserByID(id string) (any, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format, expects uuid")
	}

	condition := "users.id = ? AND users.is_verified = true AND users.is_deleted = false"

	user, err := service.repo.GetByCondition(condition, parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("no user found with id = %s", parsedUUID)
	}

	return user.ResponseObj(), nil
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
	condition__ := "users.email = ? AND users.is_deleted = false"

	user, err := service.repo.GetByCondition(condition__, email)
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

	otp_ := strings.ReplaceAll(otp, " ", "")

	if cachedOtp != otp_ {
		return nil, fmt.Errorf("the OTP you entered is incorrect. Please check and try again")
	}

	record := map[string]any{
		"is_verified": true,
	}

	condition := "users.email = ? AND users.is_verified = false AND users.is_deleted = false"

	err = service.repo.UpdateSpecificRecord(record, condition, email)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	relations := []string{"UserPassword", "Role"}
	join := "INNER JOIN user_passwords ON users.id = user_passwords.user_id"

	users, err := service.repo.FindAllByConditionWithJoin(relations, join, condition, email, true, false)
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

	isHtml := true
	subject, emailBody := helper.GetCredentialEmailFormat(users[0].FullName(), users[0].Email, users[0].UserPassword.Password, isHtml)

	go func() {
		if err := services.SmtpServer.SendEmail(users[0].Email, subject, emailBody, isHtml); err != nil {
			fmt.Printf("failed to send email to %s: %v", users[0].Email, err)
		}
	}()

	return "Your email has been successfully verified! We've sent your login credentials to your registered email address. Please check your inbox to proceed.", nil
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
	condition := "users.email = ? AND users.is_verified = false AND users.is_deleted = false"

	user, err := service.repo.GetByCondition(condition, email)
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
	isHtml := true
	subject, emailBody := helper.GetEmailVerificationFormat(user.FullName(), otp, isHtml)

	_, err = helper.SetCache(user.Email, otp, time.Duration(helper.OtpExpTime)*time.Minute)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := services.SmtpServer.SendEmail(user.Email, subject, emailBody, isHtml); err != nil {
			fmt.Printf("failed to send email to %s: %v", user.Email, err)
		}
	}()

	return "We have sent the OTP to your Email address.", nil
}

// GetUsers retrieves a list of users matching the specified query parameters.
// It filters for only verified users by default, builds a dynamic query,
// and returns a formatted list of user responses.
//
// Parameters:
//
//	queryParams (*models.UserQueryParams): The query parameters for filtering users.
//
// Returns:
//
//	[]models.UserResponse: A slice of user response objects.
//	error: An error if no data is found or if any operation fails.
func (service *Service) GetUsers(queryParams *models.UserQueryParams) ([]models.UserResponse, error) {
	queryParams.IsVerified = true

	filter := service.repo.GetFilter()

	filter = helper.BuildQuery(filter, queryParams)

	users, _, err := service.repo.FindAll(filter, "id", 0, 0)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if len(users) < 1 {
		return nil, fmt.Errorf("no data found")
	}

	return models.UserList(users).ResponseList(), nil
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
func (service *Service) AddUser(request models.UserRequest) (string, error) {
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
	err := service.repo.Create(&user)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	otp := helper.GenerateSecureOTP()
	isHtml := true
	subject, emailBody := helper.GetEmailVerificationFormat(user.FullName(), otp, isHtml)

	_, err = helper.SetCache(user.Email, otp, time.Duration(helper.OtpExpTime)*time.Minute)
	if err != nil {
		return "", err
	}

	go func() {
		if err := services.SmtpServer.SendEmail(user.Email, subject, emailBody, isHtml); err != nil {
			fmt.Printf("failed to send email to %s: %v", user.Email, err)
		}
	}()

	return "Please verify your Email Address. We have sent an OTP to the Email Address.", nil
}

// UpdateUser updates an existing user's information based on the provided user ID and new data.
// It validates the user's UUID, retrieves the current user data,
// applies the updates, and then saves the changes.
//
// Parameters:
//
//	id (string): The user's UUID in string format.
//	request (models.UserRequest): The new user data to update.
//
// Returns:
//
//	any: Typically a user response object with updated details.
//	error: An error if the update process fails.
func (service *Service) UpdateUser(id string, request models.UserRequest) (any, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format, expects uuid")
	}

	condition := "users.id = ? AND users.is_verified = ? AND users.is_deleted = ?"

	users, err := service.repo.FindAllByCondition(condition, parsedUUID, true, false)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if len(users) < 1 {
		return nil, fmt.Errorf("no user found with id = %s", parsedUUID)
	}

	user := users[0]
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.Phone = request.Phone
	user.RoleID = request.RoleID

	err = service.repo.Update(&user)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	return user.ResponseObj(), nil
}
