// Package service provides user management operations, including login, email verification,
// user retrieval, registration, and update functionalities.
// It interacts with the repository layer to perform CRUD operations and other user-related queries.
// It also handles JWT token generation and email sending for user verification.
package service

import (
	"e-commerce/models"
	"e-commerce/modules/user/dtos"
	"e-commerce/modules/user/repository"
	"e-commerce/utils/helper"
	"fmt"

	"github.com/google/uuid"
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

	condition := "users.user_id = ? AND users.is_verified = true"

	user, err := service.userRepo.GetByCondition(condition, parsedUUID)
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
func (service *Service) GetUsers(queryParams *dtos.UserQueryParams) ([]models.UserResponse, error) {
	queryParams.IsVerified = true

	filter := service.userRepo.GetFilter()

	filter = helper.BuildQuery(filter, queryParams)

	users, _, err := service.userRepo.FindAll(filter, "user_id", 0, 0)
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
func (service *Service) UpdateUser(id string, request dtos.UpdateUserRequest) (string, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("invalid id format, expects uuid")
	}

	condition := "users.user_id = ? AND users.is_verified = true"

	user, err := service.userRepo.GetByCondition(condition, parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("no user found with id = %s", parsedUUID)
	}

	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.Phone = request.Phone
	user.RoleID = request.RoleID

	err = service.userRepo.Update(user)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	return "User upadated successfully.", nil
}

// PartialUpdateUser updates specific fields of a user based on the provided patch request.
// It validates the user ID, checks if the user exists and is verified, and then applies only the
// non-empty fields from the request to the user's record.
//
// Parameters:
//   - id: a string representing the UUID of the user to update.
//   - request: PatchUserRequest struct containing optional fields to update.
//
// Returns:
//   - A success message string upon successful update.
//   - An error if the user ID is invalid, user is not found, or any DB operation fails.
func (service *Service) PartialUpdateUser(id string, request dtos.PatchUserRequest) (string, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("invalid id format, expects uuid")
	}

	condition := "users.user_id = ? AND users.is_verified = true"

	user, err := service.userRepo.GetByCondition(condition, parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("no user found with id = %s", parsedUUID)
	}

	patchData := make(map[string]any)

	if request.FirstName != "" {
		patchData["first_name"] = request.FirstName
	}

	if request.LastName != "" {
		patchData["last_name"] = request.LastName
	}

	if request.Email != "" {
		patchData["email"] = request.Email
	}

	if request.Phone != "" {
		patchData["phone"] = request.Phone
	}
	if request.RoleID != uuid.Nil {
		patchData["role_id"] = request.RoleID
	}

	err = service.userRepo.PartialUpdate(patchData, condition, parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	return "User upadated successfully.", nil
}

// DeleteUser deletes a verified user identified by the given ID from the database.
// It first validates the user ID format, checks for user existence and verification,
// and then performs a soft or hard delete based on repository logic.
//
// Parameters:
//   - id: A string representing the UUID of the user to delete.
//
// Returns:
//   - A success message string if the user is successfully deleted.
//   - An error if the ID is invalid, user does not exist, or deletion fails.
func (service *Service) DeleteUser(id string) (string, error) {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return "", fmt.Errorf("invalid id format, expects uuid")
	}

	condition := "users.user_id = ? AND users.is_verified = true"

	user, err := service.userRepo.GetByCondition(condition, parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	if user == nil {
		return "", fmt.Errorf("no user found with id = %s", parsedUUID)
	}

	err = service.userRepo.Delete(user, true)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	return "User successfully deleted", nil
}
