package services

import (
	"e-commerce/shared/models"
	"e-commerce/utils/constants"
	"fmt"
	"os"

	"gopkg.in/gomail.v2"
)

type EmailNotification models.SmtpServer

var SmtpServer *EmailNotification

func InitSmtpServer(smtpDetails models.SmtpServer) {
	SmtpServer = &EmailNotification{
		Host:     smtpDetails.Host,
		Port:     smtpDetails.Port,
		UserName: smtpDetails.UserName,
		Password: smtpDetails.Password,
	}
}

func (e *EmailNotification) SendEmail(emailTo string, subject, body string, isHTML bool) error {

	if e == nil {
		return fmt.Errorf("email service is not initialized")
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", os.Getenv(constants.EMAIL_FROM))
	mailer.SetHeader("To", emailTo)
	mailer.SetHeader("Subject", subject)

	// Set the email body type (text or HTML)
	if isHTML {
		mailer.SetBody("text/html", body)
	} else {
		mailer.SetBody("text/plain", body)
	}

	// Create a new dialer for the SMTP server
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)

	// Attempt to send the email
	if err := dialer.DialAndSend(mailer); err != nil {
		fmt.Printf("Failed to send email to %s: %v", emailTo, err)
		return fmt.Errorf("failed to send email to %v: %w", emailTo, err)
	}

	fmt.Printf("Email sent successfully to %s", emailTo)
	return nil
}
