// Package services provides email-service related functions.
// It contains functionality to initialize and send emails using SMTP.
package services

import (
	"e-commerce/models"
	"e-commerce/utils/constants"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

// EmailNotification holds SMTP server details used to send emails.
// This is an alias for models.SMTPServer but scoped to this service.
type EmailNotification models.SMTPServer

// SMTPServer is a singleton instance used across the application to send emails.
// It must be initialized before calling SendEmail.
var SMTPServer *EmailNotification

// InitSMTPServer initializes the global SMTP server instance used to send emails.
//
// Parameters:
//   - smtpDetails: A struct containing SMTP server configurations such as Host, Port, UserName, and Password.
//
// This function should be called once during application startup.
func InitSMTPServer(smtpDetails models.SMTPServer) {
	SMTPServer = &EmailNotification{
		Host:     smtpDetails.Host,
		Port:     smtpDetails.Port,
		UserName: smtpDetails.UserName,
		Password: smtpDetails.Password,
	}
}

// SendEmail sends an email using the initialized SMTP server.
//
// Parameters:
//   - emailTo: Recipient's email address.
//   - subject: Email subject line.
//   - body: The content of the email (can be plain text or HTML).
//   - isHTML: A boolean indicating whether the email body is HTML formatted.
//
// Returns:
//   - An error if the email could not be sent, otherwise nil.
//
// This method requires the SMTP server to be initialized via InitSmtpServer.
// If not initialized, it will return an error.
func (e *EmailNotification) SendEmail(
	emailTo string,
	subject,
	body string,
	isHTML bool,
) error {

	if e == nil {
		return fmt.Errorf("email service is not initialized")
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv(constants.EmailFrom))
	mailer.SetHeader("To", emailTo)
	mailer.SetHeader("Subject", subject)

	if isHTML {
		mailer.SetBody("text/html", body)
	} else {
		mailer.SetBody("text/plain", body)
	}

	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)

	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Printf("Failed to send email to %s: %v", emailTo, err)
		return fmt.Errorf("failed to send email to %v: %w", emailTo, err)
	}

	fmt.Printf("Email sent successfully to %s", emailTo)
	return nil
}
