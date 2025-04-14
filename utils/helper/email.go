// Package helper provides gmail-related utility functions for the application.
package helper

import (
	"e-commerce/utils/constants"
	"fmt"
	"os"
	"time"
)

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
