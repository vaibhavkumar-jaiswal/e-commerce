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

type Service struct {
	repo *repo.Repo
}

// NewUserService initializes a new instance of the Service struct with the provided repository.
func NewUserService(userRepo *repo.Repo) *Service {
	return &Service{
		repo: userRepo,
	}
}

// GetUserByConditionWithJoin retrieves a user by applying a join query on the `users` and `user_passwords` tables.
// It validates the user's credentials and generates a JWT token if the credentials are valid.
func (service *Service) GetUserByConditionWithJoin(data models.Login) (any, error) {
	join := "INNER JOIN user_passwords ON users.id = user_passwords.user_id"
	condition := "users.email = ? AND user_passwords.password = ? AND users.is_verified = true"

	relations := []string{"Role"}
	userList, err := service.repo.FindByConditionWithJoin(relations, join, condition, data.UserName, data.Password)
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

// GetUserByID retrieves a user by their unique ID (UUID).
// It ensures the user exists and is verified before returning their details.
func (service *Service) GetUserByID(id string) (any, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id format, expects uuid")
	}

	condition := "users.id = ? AND users.is_verified = true"

	users, err := service.repo.FindByCondition(condition, parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if len(users) < 1 {
		return nil, fmt.Errorf("no user found with id = %s", parsedUUID)
	}

	return users[0].ResponseObj(), nil
}

// VerifyEmail verifies a user's email address using an OTP.
// If the OTP is valid, the user's email is marked as verified, and their credentials are sent via email.
func (service *Service) VerifyEmail(email, otp string) (any, error) {
	condition__ := "users.email = ? AND users.is_deleted = ?"

	userList, err := service.repo.FindByCondition(condition__, email, false)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if len(userList) < 1 {
		return nil, fmt.Errorf("user with email (%s) is not registered", email)
	}

	if userList[0].IsVerified {
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

	condition := "users.email = ? AND users.is_verified = ? AND users.is_deleted = ?"

	err = service.repo.UpdateSpecificRecord(record, condition, email, false, false)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	relations := []string{"UserPassword", "Role"}
	join := "INNER JOIN user_passwords ON users.id = user_passwords.user_id"

	users, err := service.repo.FindByConditionWithJoin(relations, join, condition, email, true, false)
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

// ResendVerificationCode resends the OTP to the user's email address for email verification.
func (service *Service) ResendVerificationCode(email string) (any, error) {
	condition := "users.email = ? AND users.is_deleted = ?"

	users, err := service.repo.FindByCondition(condition, email, false)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if len(users) < 1 {
		return nil, fmt.Errorf("the user with email (%s) has not registered", email)
	}

	if users[0].IsVerified {
		return nil, fmt.Errorf("user with email (%s) is already verified, you can proceed to login", email)
	}

	otp := helper.GenerateSecureOTP()
	isHtml := true
	subject, emailBody := helper.GetEmailVerificationFormat(users[0].FullName(), otp, isHtml)

	_, err = helper.SetCache(users[0].Email, otp, time.Duration(helper.OtpExpTime)*time.Minute)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := services.SmtpServer.SendEmail(users[0].Email, subject, emailBody, isHtml); err != nil {
			fmt.Printf("failed to send email to %s: %v", users[0].Email, err)
		}
	}()

	return "We have sent the OTP to your Email address.", nil
}

// GetUsers retrieves a list of verified users based on the provided query parameters.
func (service *Service) GetUsers(queryParams *models.UserQueryParams) ([]models.UserResponse, error) {
	queryParams.IsVerified = true

	filter := service.repo.DB.Model(&models.User{})

	filter = helper.BuildQuery(filter, queryParams)

	users, _, err := service.repo.GetAll(filter, "id", 0, 0)
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

// AddUser creates a new user in the system and sends an OTP to their email for verification.
func (service *Service) AddUser(request models.UserRequest) (any, error) {
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
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
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

	return "Please verify your Email Address. We have sent an OTP to the Email Address.", nil
}
