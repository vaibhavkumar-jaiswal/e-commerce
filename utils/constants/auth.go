// Package constants provides a collection of constant values used for authentication and authorization
// related operations in the application.
package constants

// LoggedOutKey - Logged out key for Redis
const LoggedOutKey = "logged_out"

// JWT related constants
const (
	SecretKey          = "JWT_SECRET"
	UserJwtClaimKey    = "user_details"
	UserDataContextKey = "logged_in_user_data"
	UserDataOfSession  = "user_data_session"
)
