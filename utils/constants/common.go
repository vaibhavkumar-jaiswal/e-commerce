// Package constants provides a collection of constant values used throughout the application.
package constants

// file permissions
const (
	DirPerm750  = 0750 //rwx for owner, rx for group, none for others
	FilePerm640 = 0640 // rw for owner, r for group, none for others
)

// key for setting response in the gin context
const (
	ResponseDataKey   = "response_data"
	ResponseStatusKey = "response_status"
)

// RateLimitPrefix - ratelimit key prefix
const RateLimitPrefix = "rate_limit_"
